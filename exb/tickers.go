package exb

import (
	"encoding/json"
	"strings"
)

type Tickers struct {
	Pair               string
	Volume             int
	PriceChangePercent int
}

type Ticker struct {
	At     string `json:"at"`
	Ticker struct {
		Low                string      `json:"low"`
		High               string      `json:"high"`
		Open               string      `json:"open"`
		Last               string      `json:"last"`
		Volume             string      `json:"volume"`
		Amount             string      `json:"amount"`
		Vol                string      `json:"vol"`
		AvgPrice           string      `json:"avg_price"`
		PriceChangePercent string      `json:"price_change_percent"`
		At                 interface{} `json:"at"`
	} `json:"ticker"`
}

func (to *Keys) GetTicker(market string) Ticker {
	req := to.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/public/markets/"+market+"/tickers", "GET")
	var res Ticker
	_ = json.Unmarshal(req.([]uint8), &res)
	//log.Printf("Тикер: %s", res)
	return res
}
func (to *Keys) GetAllTickers(quote_asset string) []Tickers {
	req := to.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/public/markets/tickers", "GET")
	var res map[string]Ticker
	_ = json.Unmarshal(req.([]uint8), &res)
	var items []Tickers
	for pair, ticker := range res {
		if strings.Contains(pair, quote_asset) {
			items = append(items, formatTickers(pair, ticker))
		}
	}
	return items

}
