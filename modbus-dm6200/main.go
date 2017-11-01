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
	handler.SlaveId = 2
	handler.Timeout = 5 * time.Second
	err := handler.Connect()
	defer handler.Close()
	if err != nil {
		log.Println("error : ", err)
	}

	client := modbus.NewClient(handler)

	// buf := make([]byte, 16)
	// f := float32(4030201.0)
	// n := binary.PutUvarint(buf, uint64(math.Float32bits(f)))

	buf := make([]byte, 16)
	f := float32(4030201.0)
	n := binary.PutUvarint(buf, uint64(math.Float32bits(f)))

	client.WriteMultipleRegisters(320, 2, buf[:n])
	results, err := client.ReadHoldingRegisters(320, 2)
	if err != nil {
		log.Println("error : ", err)
		return
	}

	log.Println("result :", string(results), "byte : ", results)

	var x int64
	for i := 300; i < 310; i++ {
		x = int64(i)
		results, err = client.ReadHoldingRegisters(uint16(x), 1)
		if err != nil {
			log.Println("error : ", err)
			continue
		}
		log.Println("result :", string(results), "byte : ", results, "address : ", x)
		continue
	}
}
