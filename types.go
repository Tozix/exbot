package main

type (
	Exb struct {
		Key         string  `json:"key"`
		Secret      string  `json:"secret"`
		Fee         float64 `json:"fee"`
		Quote_asset string  `json:"quote_asset"`
		Coin        string  `json:"coin"`
		Sell_up     float64 `json:"sell_up"`
		Min_volume  int     `json:"min_volume"`
	}

	MySQL struct {
		Host string `json:"host_db"`
		Name string `json:"name_db"`
		User string `json:"user_db"`
		Pass string `json:"pass_db"`
	}

	// Config - настройки
	Config struct {
		Exb   Exb   `json:"Exb"`
		MySQL MySQL `json:"MySQL"`
	}
)
