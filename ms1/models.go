package main

import (
	"fmt"
	"reflect"
	"time"
)

type TransactionData struct {
	V1     float64 `json:"v1"`
	V2     float64 `json:"v2"`
	V3     float64 `json:"v3"`
	V4     float64 `json:"v4"`
	V5     float64 `json:"v5"`
	V6     float64 `json:"v6"`
	V7     float64 `json:"v7"`
	V8     float64 `json:"v8"`
	V9     float64 `json:"v9"`
	V10    float64 `json:"v10"`
	V11    float64 `json:"v11"`
	V12    float64 `json:"v12"`
	V13    float64 `json:"v13"`
	V14    float64 `json:"v14"`
	V15    float64 `json:"v15"`
	V16    float64 `json:"v16"`
	V17    float64 `json:"v17"`
	V18    float64 `json:"v18"`
	V19    float64 `json:"v19"`
	V20    float64 `json:"v20"`
	V21    float64 `json:"v21"`
	V22    float64 `json:"v22"`
	V23    float64 `json:"v23"`
	V24    float64 `json:"v24"`
	V25    float64 `json:"v25"`
	V26    float64 `json:"v26"`
	V27    float64 `json:"v27"`
	V28    float64 `json:"v28"`
	Amount float64 `json:"amount"`
}

type NormDTO struct {
	TranID string `json:"tran_id"`
	Data   [29]float64 `json:"data"`
}

type Transaction struct {
	ID        string          `json:"id"`
	Source    string          `json:"source"`
	Timestamp time.Time       `json:"timestamp"`
	Tran      TransactionData `json:"tran"`
}

func (t Transaction) String() string {
	return fmt.Sprintf("{ ID: %s, Timestamp: %s, TransactionData: %Tran }", t.ID, t.Timestamp, t.Tran)
}

func (normDTO NormDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v }", normDTO.TranID, normDTO.Data)
}

func (t Transaction) TranToNorm() NormDTO {
	values := reflect.ValueOf(t.Tran)
	data := [29]float64{}
	for i := 0; i < 29; i++ {
		data[i] = values.Field(i).Float()
	}
	return NormDTO{
		TranID: t.ID,
		Data:   data,
	}
}
