package main

import (
	"fmt"
)

type PredictionDTO struct {
	TranID  string    `json:"tran_id"`
	Data    [29]float64 `json:"data"`
	IsFraud bool      `json:"is_fraud"`
}

func (predictionDTO PredictionDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v, IsFraud: %t }", predictionDTO.TranID, predictionDTO.Data, predictionDTO.IsFraud)
}
