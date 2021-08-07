package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net"
	"os"
)

var counter = 0
var aiAddr string
var aiConn *websocket.Conn

func handleConnection(gatewayConn net.Conn, normConn *websocket.Conn) {
	defer func() {
		err := gatewayConn.Close()
		if err != nil {
			return
		}
	}()
	var output = &Transaction{}
	var err error
	decoder := json.NewDecoder(gatewayConn)

	for {
		err = decoder.Decode(output)
		if err != nil {
			if err == io.ErrUnexpectedEOF || err == io.EOF {
				break
			}
			log.Fatal(err.Error())
		}
		if err = sendOutput(output, normConn); err != nil {
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

func receiveHandler(normConn *websocket.Conn) {
	aiAddr = os.Getenv("AI_URL")
	for {
		_, msg, err := normConn.ReadMessage()
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
