package clients

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	cache2 "github.com/pitpalac36/Fraud-Detector/aggregator/cache"
	"github.com/pitpalac36/Fraud-Detector/aggregator/models"
	"os"
)

type DenormHandler struct {
	DenormConn     *websocket.Conn
	Cache          *cache2.Cache
	PredictionChan chan models.Prediction
}

func (d *DenormHandler) HandleDenormReceive() error {
	denormDTO := models.DenormDTO{}
	prediction := &models.Prediction{}
	counter := 0
	var err error
	for {
		if d.DenormConn == nil {
			d.DenormConn, _, err = websocket.DefaultDialer.Dial(os.Getenv("DENORM_URL"), nil)
			if err != nil {
				return err
			}
		}
		_, bytes, err := d.DenormConn.ReadMessage()
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &denormDTO)
		if err != nil {
			return err
		}
		prediction, err = d.Cache.Get(denormDTO.TranID)
		if err != nil {
			return err
		}
		prediction.Data = denormDTO.Data
		if prediction.IsFraud {
			counter++
			fmt.Println(counter)
			d.PredictionChan <- *prediction
		}
	}
}

func (d *DenormHandler) HandleAiReceive(predDTO *models.Prediction) error {
	var err error
	denormDTO := models.DenormDTO{}
	if d.DenormConn == nil {
		d.DenormConn, _, err = websocket.DefaultDialer.Dial(os.Getenv("DENORM_URL"), nil)
		if err != nil {
			return err
		}
	}
	err = d.Cache.Set(predDTO)
	//fmt.Println(d.Cache.Client.DBSize(d.Cache.Context).String())
	if err != nil {
		return err
	}
	denormDTO.TranID = predDTO.TranID
	denormDTO.Data = predDTO.Data
	bytes, err := json.Marshal(denormDTO)
	if err != nil {
		return err
	}
	if err = d.DenormConn.WriteJSON(bytes); err != nil {
		return err
	}
	return nil
}
