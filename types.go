package main

type (

	// Binance - настройки Binance
	Exb struct {
		Key         string  `json:"key"`
		Secret      string  `json:"secret"`
		Fee         float64 `json:"fee"`
		Quote_asset string  `json:"quote_asset"`
		Coin        string  `json:"coin"`
		Sell_up     float64 `json:"sell_up"`
		Min_volume  int     `json:"min_volume"`
	}

	// Config - настройки
	Config struct {
		Exb Exb `json:"Exb"`
	}
)
