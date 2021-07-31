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
const aiAddr = "ws://localhost:8083"
var aiConn *websocket.Conn



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
		if err = sendOutput(output, wsConn); err != nil {
			log.Fatal(err.Error())
		} else {
			counter++
			fmt.Println(counter)
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
	//normDTO := NormDTO{}
	//aiCounter := 0
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}

		if aiConn == nil {
			aiConn, _, err = websocket.DefaultDialer.Dial(aiAddr, nil)
			if err != nil {
				log.Fatal("Error connecting to Websocket Server:", err)
			}
		}
		if err = aiConn.WriteJSON(msg); err != nil {
			log.Fatal(err.Error())
		}
	}
}
