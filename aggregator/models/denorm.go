package models

import (
	"fmt"
)

type DenormDTO struct {
	TranID string      `json:"tran_id"`
	Data   [29]float64 `json:"data"`
}

func (d DenormDTO) String() string {
	return fmt.Sprintf("{ TranID: %s, Data: %v}", d.TranID, d.Data)
}
