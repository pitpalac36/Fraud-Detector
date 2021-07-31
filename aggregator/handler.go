package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const denormAddr = "ws://localhost:8085"

var denormConn *websocket.Conn

func socketHandler(w http.ResponseWriter, r *http.Request) {
	aiConn, err := upgrader.Upgrade(w, r, nil)
	predDTO := PredictionDTO{}
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
		_, bytes, err := aiConn.ReadMessage()
		if err != nil {
			log.Fatal(err.Error())
		}
		err = json.Unmarshal(bytes, &predDTO)
		if err != nil {
			log.Fatal(err.Error())
		}
		err = handleAiData(&predDTO)
		if err != nil {
			return
		}
	}
}

func handleAiData(predDTO *PredictionDTO) error {
	var err error
	denormDTO := DenormDTO{}
	if denormConn == nil {
		denormConn, _, err = websocket.DefaultDialer.Dial(denormAddr, nil)
		if err != nil {
			return err
		}
	}
	denormDTO.TranID = predDTO.TranID
	denormDTO.Data = predDTO.Data
	fmt.Println(denormDTO)
	bytes, err := json.Marshal(denormDTO)
	if err != nil {
		return err
	}
	if err = denormConn.WriteJSON(bytes); err != nil {
		return err
	}
	return nil
}
