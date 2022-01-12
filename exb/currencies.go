package exb

import (
	"encoding/json"
)

type Currencies struct {
	Group      string `json:"group"`
	MarketId   string `json:"market_id"`
	MarketType string `json:"market_type"`
	Limit      string `json:"limit"`
	Page       string `json:"page"`
	Ordering   string `json:"ordering"`
	Order_by   string `json:"order_by"`
}

func GetCurrencies() Fees {
	req := PublicRequest("https://www.exbitron.com", "/api/v2/peatio/public/trading_fees")
	var res Fees
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
