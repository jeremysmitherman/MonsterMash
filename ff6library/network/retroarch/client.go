package retroarch

import (
	"MonsterMash/ff6library/network"
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

func (c *Client) ListenAndServe(dataChannel chan []uint16, exitRequested chan interface{}, finished chan interface{}) {
	c.Status = "Waiting for Emulator..."
	log.Println(c.Status)
	conn := new(net.UDPConn)

	// Not sure if needed for a UDP connection, but I'm not going to be the reason someone's chromebook dies
	// Close the 'finished' channel or the program will never exit
	defer func() {
		c.Status = "Shutting Down"
		fmt.Println("Shutting down EMU listener...")
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

	for {
		select {
		case <-exitRequested:
			// Something external closed this channel, return and let the defer clean up
			fmt.Println("Listener Exit requested")
			return
		default:
			time.Sleep(1 * time.Second)
			_, err = conn.Write([]byte("READ_CORE_RAM 2001 12"))
			if err != nil {
				c.ConnectionState = network.ERROR
				c.Status = "Error Writing to EMU"
				c.StatusMessage = err.Error()
				log.Println(c.Status + " " + c.StatusMessage)
				break
			}

			_ = conn.SetReadDeadline(time.Now().Add(time.Second * 2))
			data, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				c.ConnectionState = network.ERROR
				c.Status = "No emulator connection"
				c.StatusMessage = err.Error()
				log.Println(c.Status + " " + c.StatusMessage)
				break
			}

			// If we got this far, everything is obviously talking to each other, inform the frontend
			// If c.Status isn't "Connected" the frontend won't show anything
			if c.Status != "Connected" {
				log.Println("Connected to EMU")
			}
			c.ConnectionState = network.CONNECTED
			c.Status = "Connected"
			parsedData := ParseRetroArchReturn(data)

			// If for whatever reason we can't write to the data channel, discard the message
			select {
			case dataChannel <- parsedData:
				log.Printf("sending %v\n", parsedData)
			default:
				log.Println("Failed to send, data channel blocked?")
			}
		}
	}
}

func ParseRetroArchReturn(encounterString string) []uint16 {
	// RA returns a fairly verbose response over the wire, so remove the extra info and spaces
	encounterString = strings.Replace(
		strings.Replace(encounterString, "READ_CORE_RAM 2001 ", "", -1),
		" ",
		"",
		-1)

	var monsters []uint16
	if len(encounterString) < 24 {
		return monsters
	}

	for idx := 0; idx < 24; idx += 4 {
		// RAM values are in the opposite endianess we expect, so split it out so we can flip it later
		monsterIdxString := encounterString[idx : idx+4]
		msb := monsterIdxString[2:4]
		lsb := monsterIdxString[0:2]

		monsterIdx, err := strconv.ParseUint(msb+lsb, 16, 16)
		if err != nil {
			log.Println(err)
			continue
		}

		// There's no valid monsters above this value, but there's non-related reasons this value could be higher
		if monsterIdx > 383 {
			continue
		}

		// Don't add if we already have this monster index
		add := true
		for _, m := range monsters {
			if m == uint16(monsterIdx) {
				add = false
				break
			}
		}

		if add {
			monsters = append(monsters, uint16(monsterIdx))
		}
	}

	return monsters
}
