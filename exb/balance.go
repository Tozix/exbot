package exb

import (
	"encoding/json"
)

type Balance struct {
	Balance string `json:"balance"`
	Locked  string `json:"locked"`
}

func (to *Keys) Balance(currency string) Balance {

	req := to.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/account/balances/"+currency, "GET")
	var res Balance
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
