package main

import (
	"io/ioutil"
	"sync"
)

func handleData(request *Request) error {
	block := request.buffer[2:4]
	rm := sync.Mutex{}
	rm.Lock()
	filename := fd[request.addr.Port]
	rm.Unlock()

	err := ioutil.WriteFile(filename, request.buffer[4:request.size], 0644)
	if err != nil {
		return err
	}

	if request.size-6 < 512 {
		rm.Lock()
		delete(fd, request.addr.Port)
		rm.Unlock()
	}

	request.sendAck(block)
	return nil
}
