package models

import "fmt"

type Prediction struct {
	TranID  string      `json:"tran_id"`
	Data    [29]float64 `json:"data"`
	IsFraud bool        `json:"is_fraud"`
}

func (p Prediction) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v, IsFraud: %t }", p.TranID, p.Data, p.IsFraud)
}

func (p Prediction) ToResponse() *Response {
	return &Response{
		Source:    p.TranID,
		Timestamp: p.TranID,
		Amount:    p.Data[28],
		V1:        p.Data[0],
		V2:        p.Data[1],
		V3:        p.Data[2],
		V4:        p.Data[3],
		V5:        p.Data[4],
		V6:        p.Data[5],
		V7:        p.Data[6],
		V8:        p.Data[7],
		V9:        p.Data[8],
		V10:       p.Data[9],
		V11:       p.Data[10],
		V12:       p.Data[11],
		V13:       p.Data[12],
		V14:       p.Data[13],
		V15:       p.Data[14],
		V16:       p.Data[15],
		V17:       p.Data[16],
		V18:       p.Data[17],
		V19:       p.Data[18],
		V20:       p.Data[19],
		V21:       p.Data[20],
		V22:       p.Data[21],
		V23:       p.Data[22],
		V24:       p.Data[23],
		V25:       p.Data[24],
		V26:       p.Data[25],
		V27:       p.Data[26],
		V28:       p.Data[27],
	}
}