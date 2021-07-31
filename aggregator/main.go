package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
/*
   TODO: wait for denormalized data (ws)
   TODO: send data to dashboard
*/

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", socketHandler)
	if err:= http.ListenAndServe("localhost:8084", nil); err != nil {
		log.Fatal(err.Error())
	}
}