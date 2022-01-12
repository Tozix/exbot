package exb

import (
	"encoding/json"
	"time"
)

type Fees []struct {
	ID         int       `json:"id"`
	Group      string    `json:"group"`
	MarketID   string    `json:"market_id"`
	MarketType string    `json:"market_type"`
	Maker      string    `json:"maker"`
	Taker      string    `json:"taker"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func GetFees() Fees {
	req := PublicRequest("https://www.exbitron.com", "/api/v2/peatio/public/trading_fees")
	var res Fees
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
