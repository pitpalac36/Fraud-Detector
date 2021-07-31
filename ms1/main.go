package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net"
)

var ln net.Listener
var wsConn *websocket.Conn

func init() {
	var err error
	ln, err = net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	socketUrl := "ws://localhost:8082"
	wsConn, _, err = websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
}

func main() {
	defer func() {
		err := wsConn.Close()
		fmt.Println("Closing normalizer web socket")
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(conn, wsConn)
		go receiveHandler(wsConn)
	}
}
