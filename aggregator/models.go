package main

import (
	"fmt"
)

type PredictionDTO struct {
	TranID  string      `json:"tran_id"`
	Data    [29]float64 `json:"data"`
	IsFraud bool        `json:"is_fraud"`
}

func (p PredictionDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v, IsFraud: %t }", p.TranID, p.Data, p.IsFraud)
}

type DenormDTO struct {
	TranID string      `json:"tran_id"`
	Data   [29]float64 `json:"data"`
}

func (d DenormDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v}", d.TranID, d.Data)
}
