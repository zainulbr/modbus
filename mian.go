package main

import (
	"log"
	"time"

	"github.com/goburrow/modbus"
)

func main() {

	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		log.Println("error : ", err)
	}

	client := modbus.NewClient(handler)
	readRegister(3925, 2, client)

}

func readRegister(address, quantity uint16, client modbus.Client) {
	results, err := client.ReadHoldingRegisters(address-1, quantity)
	if err != nil {
		log.Println("error : ", err)
		return
	}
	log.Println("result :", string(results), "byte : ", results)
}
