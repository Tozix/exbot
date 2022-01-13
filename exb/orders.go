package exb

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

//Список всех сделок
func (ex *Keys) GetOrders() Orders {
	req := ex.PrivateRequest("https://www.exbitron.com", nil, "/api/v2/peatio/market/orders")
	var res Orders
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}

// GetOpenOrders - получение списка открытых ордеров
func (ex *Keys) GetOpenOrders(market string) ([]Order, error) {
	openOrders, err := ex.PR("https://www.exbitron.com", nil, "/api/v2/peatio/market/orders")
	if err != nil {
		return nil, errors.New("не удалось получить открытые ордера: " + err.Error())
	}

	/*
		openOrders, err := api.Client.NewListOpenOrdersService().Symbol(pair).Do(context.Background())
		if err != nil {
			return nil, errors.New("не удалось получить открытые ордера: " + err.Error())
		}
	*/
	//log.Printf("RESA: %v", openOrders)

	for _, order := range openOrders {
		//if order.State == "wait" {
		log.Printf("IDы: %v Магаз3333333: %v Объеб монет: %v", order.ID, order.Market, order.OriginVolume)
		//}
	}

	formattedOpenOrders := make([]Order, len(openOrders))

	for index, order := range openOrders {
		formattedOpenOrders[index] = *order
		//formattedOpenOrders[index] = formatOrder(*order)
	}

	return formattedOpenOrders, nil
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
	params := map[string]string{
		"id": ID,
	}
	req := ex.PrivateRequest("https://www.exbitron.com", params, "/api/v2/peatio/market/orders/"+ID+"/cancel")
	var res Order
	_ = json.Unmarshal(req.([]uint8), &res)
	return res
}
