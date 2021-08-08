package handlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/pitpalac36/Fraud-Detector/aggregator/clients"
	"github.com/pitpalac36/Fraud-Detector/aggregator/models"
	"log"
	"net/http"
)

type AIHandler struct {
	AiConn        *websocket.Conn
	DenormHandler *clients.DenormHandler
}

func (h *AIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	aiConn, err := upgrader.Upgrade(w, r, nil)
	predDTO := models.Prediction{}
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
	go func() {
		err := h.DenormHandler.HandleDenormReceive()
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
		err = h.DenormHandler.HandleAiReceive(&predDTO)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

