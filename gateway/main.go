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
	if os.Getenv("PRODUCTION") != "1" {
		err = godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	ln, err = net.Listen("tcp", ":8080")
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
