package main

import (
	"fmt"

	"github.com/goburrow/modbus"
)

func main() {
	// handler := modbus.NewTCPClientHandler("localhost:1502")
	handler := modbus.NewTCPClientHandler("192.168.31.98:1502")
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	_, err = client.WriteMultipleRegisters(0, 3, []byte{0, 5, 0, 4, 0, 5})
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	results, err := client.ReadHoldingRegisters(0, 3)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("results %v\n", results)
}
