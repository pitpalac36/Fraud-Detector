package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	counter := 0
	var predictionDTO = &PredictionDTO{}
	var err error
	decoder := json.NewDecoder(conn)
	for {
		err = decoder.Decode(predictionDTO)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		counter++
		fmt.Println(counter)
	}
}
