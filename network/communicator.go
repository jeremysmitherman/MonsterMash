package network

import "os"

type ConnectionState uint8

const (
	IDLE ConnectionState = iota
	CONNECTING
	CONNECTED
	CLOSED
	ERROR
)

type Communicator interface {
	Listen(chan os.Signal)
}
