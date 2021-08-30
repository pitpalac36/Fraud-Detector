package handlers

import (
	"github.com/gorilla/websocket"
	"github.com/pitpalac36/Fraud-Detector/aggregator/models"
	"log"
	"net/http"
)

type UIHandler struct {
	Conn           *websocket.Conn
	PredictionChan chan models.Prediction
}

func (h *UIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	u := websocket.Upgrader{}
	u.CheckOrigin = func(r *http.Request) bool { return true }
	var err error
	h.Conn, err = u.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	defer func() {
		err = h.Conn.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()

	var res models.Prediction
	for {
		select {
			case res = <- h.PredictionChan:
				//fmt.Println(*res.ToResponse())
				if err = h.Conn.WriteJSON(res.ToResponse()); err != nil {
					log.Fatal(err)
				}
		}

	}
}
