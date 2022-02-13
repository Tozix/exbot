package exb

import (
	"encoding/json"
	"fmt"
	"time"
)

type Orders []struct {
	Order
}
type Order struct {
	ID              int       `json:"id"`
	UUID            string    `json:"uuid"`
	Side            string    `json:"side"`
	OrdType         string    `json:"ord_type"`
	Price           string    `json:"price"`
	AvgPrice        string    `json:"avg_price"`
	State           string    `json:"state"`
	Market          string    `json:"market"`
	MarketType      string    `json:"market_type"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	OriginVolume    string    `json:"origin_volume"`
	RemainingVolume string    `json:"remaining_volume"`
	ExecutedVolume  string    `json:"executed_volume"`
	MakerFee        string    `json:"maker_fee"`
	TakerFee        string    `json:"taker_fee"`
	TradesCount     int       `json:"trades_count"`
}

// GetOpenOrders - получение списка открытых ордеров
func (ex *Keys) GetOpenOrders(market string) ([]Order, error) {
	params := map[string]string{
		"market": market,
		"state":  "wait",
	}
	req := ex.PrivateRequest("https://www.exbitron.com", params, "/api/v2/peatio/market/orders", "Любая хуйня если хотим отправить ГЕТ")
	var res []*Order
	_ = json.Unmarshal(req.([]uint8), &res)
	formattedOpenOrders := make([]Order, len(res))
	for index, order := range res {
		formattedOpenOrders[index] = *order
	}

	return formattedOpenOrders, nil
}
func (ex *Keys) GetOrder(ID string) Order {
	req := ex.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/market/orders/"+ID, "GET")
	var res Order
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}

//Разместить ордер на продажу
func (ex *Keys) SellOrder(market string, volume float64, price float64) Order {
	params := map[string]string{
		"market":   market,
		"volume":   fmt.Sprintf("%.8f", volume),
		"price":    fmt.Sprintf("%.8f", price),
		"side":     "sell",
		"ord_type": "limit",
	}
	req := ex.PrivateRequest("https://www.exbitron.com", params, "/api/v2/peatio/market/orders")
	var res Order
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}

//Убрать все сделки (по параметрам)
func (ex *Keys) CancelOrders(market string) Orders {
	params := map[string]string{
		"market": market,
		"side":   "sell",
	}
	req := ex.PrivateRequest("https://www.exbitron.com", params, "/api/v2/peatio/market/orders/cancel")
	var res Orders
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}

//Убрать сделку по ID
func (ex *Keys) CancelOrder(ID string) Order {
	req := ex.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/market/orders/"+ID+"/cancel", "GET")
	var res Order
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
