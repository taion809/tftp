package main

import (
	"os"
	"sync"
)

func handleWrq(request *Request) error {
	filename, _ := parseWrqBuffer(request)

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	file.Close()

	rm := sync.Mutex{}
	rm.Lock()
	fd[request.addr.Port] = filename
	rm.Unlock()

	request.sendAck([]byte{00, 00})
	return nil
}

func parseWrqBuffer(request *Request) (string, string) {
	lastByte := 2

	filename := ""
	mode := ""

	for _, v := range request.buffer[lastByte:] {
		lastByte++
		if v == 0 {
			break
		}

		filename += string(v)
	}

	for _, v := range request.buffer[lastByte:] {
		lastByte++
		if v == 0 {
			break
		}

		mode += string(v)
	}

	return filename, mode
}
