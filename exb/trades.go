package exb

import (
	"encoding/json"
	"time"
)

type Trades []struct {
	ID          int       `json:"id"`
	Price       string    `json:"price"`
	Amount      string    `json:"amount"`
	Total       string    `json:"total"`
	FeeCurrency string    `json:"fee_currency"`
	Fee         string    `json:"fee"`
	FeeAmount   string    `json:"fee_amount"`
	Market      string    `json:"market"`
	MarketType  string    `json:"market_type"`
	CreatedAt   time.Time `json:"created_at"`
	TakerType   string    `json:"taker_type"`
	Side        string    `json:"side"`
	OrderID     int       `json:"order_id"`
}

func (ex *Keys) Trades() Trades {

	req := ex.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/market/trades")
	var res Trades
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
