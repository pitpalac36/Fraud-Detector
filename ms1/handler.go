package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
)

var counter = 0

func handleConnection(conn net.Conn, wsConn *websocket.Conn) {
	var output = &Transaction{}
	var err error
	decoder := json.NewDecoder(conn)

	for {
		err = decoder.Decode(output)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		counter++
		//fmt.Println(output)

		if err = sendOutput(output, wsConn); err != nil {
			log.Fatal(err.Error())
		}
	}
}

func sendOutput(output *Transaction, wsConn *websocket.Conn) error {
	normDTO := output.TranToNorm()
	bytes, err := json.Marshal(normDTO)
	if err != nil {
		return err
	}
	if err := wsConn.WriteJSON(bytes); err != nil {
		return err
	}
	return nil
}

func receiveHandler(connection *websocket.Conn) {
	for {
		_, msg, err := connection.ReadMessage()
		fmt.Println(msg)
		if err != nil {
			log.Fatal("Error in receive:", err.Error())
			return
		}
		log.Printf("Received: %s\n", msg)
	}
}
