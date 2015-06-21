package main

import (
	"bytes"
	"log"
	"net"
)

var (
	RRQ  = []byte{00, 01}
	WRQ  = []byte{00, 02}
	DATA = []byte{00, 03}
	ACK  = []byte{00, 04}
)

type Request struct {
	buffer []byte
	addr   *net.UDPAddr
	size   int
}

func (r *Request) sendAck(block []byte) (int, error) {
	return connection.WriteToUDP([]byte{ACK[0], ACK[1], block[0], block[1]}, r.addr)
}

func (r *Request) Opcode() []byte {
	return r.buffer[0:2]
}

func dispatch(request *Request) {
	var err error
	opcode := request.Opcode()
	switch {
	case bytes.Equal(opcode, RRQ):
	case bytes.Equal(opcode, WRQ):
		err = handleWrq(request)
	case bytes.Equal(opcode, DATA):
		err = handleData(request)
	}

	if err != nil {
		log.Fatal(err)
	}
}
