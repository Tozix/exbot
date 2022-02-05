package exb

/*
type Ticker []struct {
	TickerEntry
}
type TickerEntry struct {
	BaseID      int    `json:"base_id"`
	QuoteID     int    `json:"quote_id"`
	LastPrice   string `json:"last_price"`
	BaseVolume  string `json:"base_volume"`
	QuoteVolume string `json:"quote_volume"`
	IsFrozen    int    `json:"isFrozen"`
}

func (to *Keys) Ticker() Ticker {

	req := to.FirstR("https://www.exbitron.com", nil, "/api/v2/peatio/coinmarketcap/ticker", "GET")
	var res Ticker
	_ = json.Unmarshal(req.([]uint8), &res)
	log.Printf("Тикер: %s", res)
	return res
}
*/
