package retroarch

import (
	"MonsterMash/network"
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

var CommandAddr string = "127.0.0.1:55355"

type Client struct {
	ConnectionState network.ConnectionState
	Status          string
	StatusMessage   string
}

func (c *Client) ListenAndServe(encounter chan uint16, exitRequested chan interface{}, finished chan interface{}) {
	c.Status = "Waiting for Emulator..."
	log.Println(c.Status)
	conn := new(net.UDPConn)

	defer func() {
		c.Status = "Shutting Down"
		fmt.Println("Shutting down listener...")
		err := conn.Close()
		if err != nil {
			close(finished)
		}
		close(finished)
	}()

	udpAddr, err := net.ResolveUDPAddr("udp", CommandAddr)
	if err != nil {
		log.Println(err)
	}

	conn, err = net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Println(err)
	}

	log.Println("UDP dialer set up")
	for {
		select {
		case <-exitRequested:
			fmt.Println("Listener Exit requested")
			return
		default:
			time.Sleep(1 * time.Second)
			log.Println("Requesting RAM...")
			_, err = conn.Write([]byte("READ_CORE_RAM 11E0 1"))
			if err != nil {
				c.ConnectionState = network.ERROR
				c.Status = "Error Writing to EMU"
				c.StatusMessage = err.Error()
				break
			}

			log.Println("Reading RAM...")
			_ = conn.SetReadDeadline(time.Now().Add(time.Second * 2))

			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				c.ConnectionState = network.ERROR
				c.Status = "No emulator connection"
				c.StatusMessage = err.Error()
				break
			}

			log.Println("Got " + data)
			encounterString := strings.TrimSpace(strings.Replace(data, "READ_CORE_RAM 11e0 ", "", -1))
			encounterIndex, err := strconv.ParseInt(encounterString, 16, 16)
			if err != nil {
				c.Status = "Error Parsing Response from EMU"
				c.StatusMessage = err.Error()
				log.Println(err)
				continue
			}

			if encounterIndex < 0 {
				c.Status = "Waiting for encounter..."
				continue
			}

			c.ConnectionState = network.CONNECTED
			c.Status = "Connected"

			select {
			case encounter <- uint16(encounterIndex):
			default:
			}
		}
	}
}
