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
	defer func() {
		err := conn.Close()
		if err != nil {
			return
		}
	}()
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
		fmt.Println(counter)

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
	if err = wsConn.WriteJSON(bytes); err != nil {
		return err
	}
	return nil
}

func receiveHandler(connection *websocket.Conn) {
	var normDTO NormDTO
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		err = json.Unmarshal(msg, &normDTO)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		//fmt.Println(normDTO)
	}
}
