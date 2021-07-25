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
const aiAddr = "localhost:8083"
var ai_conn net.Conn = nil


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
	//normDTO := NormDTO{}
	aiCounter := 0
	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			break
		}
		//err = json.Unmarshal(msg, &normDTO)
		//if err != nil {
		//	if err == io.EOF {
		//		break
		//	}
		//	log.Fatal(err.Error())
		//}
		//fmt.Println(normDTO)

		if ai_conn == nil {
			fmt.Println("Creating new tcp connection to AI module")
			ai_conn, err = net.Dial("tcp", aiAddr)
			if err != nil {
				log.Fatal(err.Error())
				return
			}
		}
		if _, err = ai_conn.Write(msg); err != nil {
			log.Fatal(err.Error())
			return
		}
		aiCounter++
		fmt.Println(aiCounter)
	}
}
