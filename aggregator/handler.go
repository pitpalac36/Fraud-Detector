package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
)

var denormAddr string

var denormConn *websocket.Conn

type AIHandler struct {
	aiConn        *websocket.Conn
	denormHandler *DenormHandler
}

func (h *AIHandler) handle(w http.ResponseWriter, r *http.Request) {
	denormAddr = os.Getenv("DENORM_URL")
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
	go func() {
		err := h.denormHandler.handleDenormReceive()
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
		err = h.handleAiReceive(&predDTO)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}

func (h *AIHandler) handleAiReceive(predDTO *PredictionDTO) error {
	var err error
	denormDTO := DenormDTO{}
	if denormConn == nil {
		denormConn, _, err = websocket.DefaultDialer.Dial(denormAddr, nil)
		if err != nil {
			return err
		}
	}
	err = h.denormHandler.cache.Set(predDTO)
	fmt.Println(h.denormHandler.cache.Client.DBSize(h.denormHandler.cache.Context).String())
	if err != nil {
		return err
	}
	denormDTO.TranID = predDTO.TranID
	denormDTO.Data = predDTO.Data
	bytes, err := json.Marshal(denormDTO)
	if err != nil {
		return err
	}
	if err = denormConn.WriteJSON(bytes); err != nil {
		return err
	}
	return nil
}

type DenormHandler struct {
	denormConn *websocket.Conn
	cache      *Cache
}

func (d *DenormHandler) handleDenormReceive() error {
	denormDTO := DenormDTO{}
	predDTO := &PredictionDTO{}
	counter := 0
	var err error
	for {
		if denormConn == nil {
			denormConn, _, err = websocket.DefaultDialer.Dial(denormAddr, nil)
			if err != nil {
				return err
			}
		}
		_, bytes, err := denormConn.ReadMessage()
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &denormDTO)
		if err != nil {
			return err
		}
		predDTO, err = d.cache.Get(denormDTO.TranID)
		if err != nil {
			return err
		}
		predDTO.Data = denormDTO.Data
		counter++
		fmt.Println(counter)
		if predDTO.IsFraud {
			fmt.Println(predDTO)
		}
	}
	// send preDTO to react

}
