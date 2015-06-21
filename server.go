package main

import (
	"log"
	"net"
)

var connection *net.UDPConn
var fd map[int]string

func init() {
	fd = make(map[int]string)
}

func main() {
	var err error

	log.Println("Starting server on 0.0.0.0:5555")

	address := &net.UDPAddr{IP: net.IP{0, 0, 0, 0}, Port: 5555}
	connection, err = net.ListenUDP("udp", address)
	connection.SetReadBuffer(1048576)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	for {
		payload := make([]byte, 1024)
		read, addr, err := connection.ReadFromUDP(payload)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Request: read=%d addr=%s\n", read, addr)
		go dispatch(&Request{buffer: payload, addr: addr, size: read})
	}

}
