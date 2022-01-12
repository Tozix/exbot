package exb

import (
	"encoding/json"
	"log"
)

func (ex *Keys) Pnl(pnl_currency string) Balance {
	log.Println(pnl_currency)
	req := ex.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/account/stats/pnl")
	var res Balance
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
