package exb

import (
	"encoding/json"
)

type Tickers struct {
	Market string
	Ticker
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
func (to *Keys) GetAllTickers() {

	req := to.FirstR("https://www.exbitron.com", nil, "/api/v2/peatio/public/markets/tickers", "GET")

	//var res Tickers
	var res map[string]Ticker
	var tickers Tickers
	_ = json.Unmarshal(req.([]uint8), &res)
	for market, ticker := range res {
		//if order.State == "wait" {
		tickers.Market = market
		tickers.Ticker = ticker
		//log.Printf("Market: %v Value: %v", market, ticker.Ticker.High)
		//	}
	}

	/*
		var res []*Ticker
		_ = json.Unmarshal(req.([]uint8), &res)
		formatTickers := make([]Ticker, len(res))
		log.Println(res)
		for index, ticker := range res {
			formatTickers[index] = *ticker
			log.Printf("Индекс: %v Значение: %v", index, ticker)
		}
	*/

	//return formatTickers
}
