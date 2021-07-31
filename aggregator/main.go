package main

/* TODO: wait for data from AI 		(tcp)
   TODO: send data to denormalizer  (ws)
   TODO: wait for denormalized data (ws)
   TODO: aggregate data
   TODO: send data to dashboard
*/

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)
var receiveAiConn *websocket.Conn
var upgrader = websocket.Upgrader{}

func init() {
	var err error
	// client (wait from denormalizer)
	receiveAiUrl := "ws://localhost:8083"
	receiveAiConn, _, err = websocket.DefaultDialer.Dial(receiveAiUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server:", err)
	}
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	defer func() {
		err := receiveAiConn.Close()
		fmt.Println("Closing ai receive web socket")
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	for {
		_, msg, err := receiveAiConn.ReadMessage()
		if err != nil {
			log.Fatal(err.Error())
		}
		go receiveHandler(msg)
	}
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	aiSendConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer func(aiConn *websocket.Conn) {
		err := aiConn.Close()
		if err != nil {

		}
	}(aiSendConn)
}
