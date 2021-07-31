package main

import (
	"fmt"
	"log"
	"net/http"
)

func receiveAiData(msg []byte) {
	fmt.Println("received from ai")
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	aiConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation:", err)
		return
	}
	defer func() {
		err := aiConn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	for {
		_, predDTO, err := aiConn.ReadMessage()
		if err != nil {
			log.Fatal(err.Error())
		}
		handleAiData(predDTO)
	}
}

func handleAiData(dto []byte) {
	fmt.Println("handling ai data!")
}
