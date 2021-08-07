package main

import (
	"github.com/joho/godotenv"
	"log"
	"net"
	"os"
)

var ln net.Listener

func init() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ln, err = net.Listen("tcp", os.Getenv("GATEWAY_ADDR"))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(conn)
	}
}
