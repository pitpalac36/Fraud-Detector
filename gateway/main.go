package main

import (
	"log"
	"net"
)

var ln net.Listener

func init() {
	var err error
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
