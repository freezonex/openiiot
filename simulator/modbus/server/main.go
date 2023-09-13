package main

import (
	"log"
	"time"

	"freezonex/openiiot/simulator/modbus/mbserver"
)

func main() {
	serv := mbserver.NewServer()
	err := serv.ListenTCP("192.168.31.98:1502")
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
