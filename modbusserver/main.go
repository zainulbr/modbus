package main

import (
	"log"
	"time"

	"github.com/goburrow/serial"
	"github.com/tbrandon/mbserver"
)

func main() {
	serv := mbserver.NewServer()
	err := serv.ListenTCP("127.0.0.1:1502")
	if err != nil {
		log.Printf("%v\n", err)
	}

	err = serv.ListenTCP("0.0.0.0:3502")
	if err != nil {
		log.Printf("%v\n", err)
	}
	client := modbus.NewClient(handler)
	
	err = serv.ListenRTU(&serial.Config{
		Address:  "/dev/ttyUSB0",
		BaudRate: 19200,
		DataBits: 8,
		StopBits: 1,
		Parity:   "E",
		Timeout:  10 * time.Second})
	if err != nil {
		log.Fatalf("failed to listen, got %v\n", err)
	}

 serv.

}
