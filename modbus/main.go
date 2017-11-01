package main

import (
	"encoding/binary"
	"log"
	"math"
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

	// buf := make([]byte, 4)
	// f := 4030201.0
	// binary.BigEndian.PutUint32(buf, math.Float32bits(float32(f)))
	// client.WriteMultipleRegisters(320, 2, buf[:])
	// results, err := client.ReadHoldingRegisters(320, 2)
	// if err != nil {
	// 	log.Println("error : ", err)
	// 	return
	// }
	
	buf := make([]byte, 16)
	f := 4030201.0
	n := binary.PutUvarint(buf, uint64(math.Float64bits(f)))
	client.WriteMultipleRegisters(320, 2, buf[:n])
	buf := make([]byte, 4)
	f := 4030201.0
	binary.PutUvarint(buf, math.Float64bits(f))
	client.WriteMultipleRegisters(320, 2, buf[:])
	results, err := client.ReadHoldingRegisters(320, 2)
	if err != nil {
		log.Println("error : ", err)
		return
	}

	log.Println("result :", string(results), "byte : ", results)
	results, err = client.ReadHoldingRegisters(3907, 2)
	if err != nil {
		log.Println("error : ", err)
		return
	}
	log.Println("result :", string(results), "byte : ", results)

}
