package main

import (
	"fmt"
)

type PredictionDTO struct {
	TranID  string    `json:"tran_id"`
	Data    [29]float64 `json:"data"`
	IsFraud bool      `json:"is_fraud"`
}

type DenormDTO struct {
	TranID  string    `json:"tran_id"`
	Data    [29]float64 `json:"data"`
}

func (predictionDTO PredictionDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v, IsFraud: %t }", predictionDTO.TranID, predictionDTO.Data, predictionDTO.IsFraud)
}

func (denormDTO DenormDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v}", denormDTO.TranID, denormDTO.Data)
}