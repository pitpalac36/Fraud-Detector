package main

/* TODO: wait for data from AI 		(tcp)
   TODO: send data to denormalizer  (ws)
   TODO: wait for denormalized data (ws)
   TODO: aggregate data
   TODO: send data to dashboard
*/

import (
	"fmt"
	"log"
	"net"
)

var ln net.Listener

func init() {
	var err error
	ln, err = net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	fmt.Println("I'm in")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConnection(conn)
	}
}
