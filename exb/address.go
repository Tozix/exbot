package exb

import (
	"encoding/json"
)

func (to *Keys) GetAddress(iso string) Balance {
	/*
		params := map[string]string{
			"iso": iso,
			"new": "0",
		}
	*/
	req := to.PrivateRequest("https://api.dex-trade.com", nil, "/v1/private/balances")
	var res Balance
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
