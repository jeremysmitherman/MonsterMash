package network

type ConnectionState uint8

const (
	IDLE ConnectionState = iota
	CONNECTING
	CONNECTED
	CLOSED
	ERROR
)

type Communicator interface {
	ListenAndServe(dataChannel chan []uint16, exitRequested chan interface{}, finished chan interface{})
}
