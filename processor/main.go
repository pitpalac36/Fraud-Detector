package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

var ln net.Listener
var normConn *websocket.Conn

func init() {
	var err error
	log.Println("It works!")
	if os.Getenv("PRODUCTION") != "1" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ln, err = net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err.Error())
	}

	normConn, _, err = websocket.DefaultDialer.Dial(os.Getenv("NORM_URL"), nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
}

func main() {
	defer func() {
		err := normConn.Close()
		fmt.Println("Closing normalizer web socket")
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	for {
		gatewayConn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(gatewayConn, normConn)
		go receiveHandler(normConn)
	}
}
