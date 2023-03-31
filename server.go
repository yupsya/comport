package main

import (
	"serial"
	"log"
	"fmt"
	"time"
)

func main() {
	config := &serial.Config{Name: "COM2", Baud: 115200, ReadTimeout: time.Millisecond * 10}
	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
		}
	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil{
			log.Fatal(err)
		}
		s := string(buf[:n])
		fmt.Print(s)
		stream.Write(buf[:n])
	}
}
