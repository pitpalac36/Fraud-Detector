package main

/* TODO: wait for data from AI 		(tcp)
   TODO: send data to denormalizer  (ws)
   TODO: wait for denormalized data (ws)
   TODO: aggregate data
   TODO: send data to dashboard
*/

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func main() {
	http.HandleFunc("/", socketHandler)
	if err:= http.ListenAndServe("localhost:8084", nil); err != nil {
		log.Fatal(err.Error())
	}
}