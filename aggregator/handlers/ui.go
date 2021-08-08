package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/pitpalac36/Fraud-Detector/aggregator/models"
	"log"
	"net/http"
)

type UIHandler struct {
	Conn           *websocket.Conn
	PredictionChan *chan models.Prediction
}

func (h *UIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{}
	var err error
	h.Conn, err = u.Upgrade(w, r, nil)

	defer func() {
		err = h.Conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	for {
		fmt.Println(<-*h.PredictionChan)
		//if err = h.Conn.WriteJSON(res.ToResponse()); err != nil {
		//	log.Fatal(err)
		//}
	}
}
