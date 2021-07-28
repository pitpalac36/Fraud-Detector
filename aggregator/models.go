package main

import (
	"fmt"
)

type PredictionDTO struct {
	TranID  string `json:"tran_id"`
	IsFraud bool   `json:"is_fraud"`
}

func (predictionDTO PredictionDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, IsFraud: %t }", predictionDTO.TranID, predictionDTO.IsFraud)
}
