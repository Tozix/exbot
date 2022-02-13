package exb

import (
	"log"
	"strconv"
	"strings"
)

func StringToFloat(str string) float64 {
	str = strings.ReplaceAll(str, ",", "")
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Printf("%f of type %T", f, f)
	}
	return f
}

func StringToInit(str string) int {
	str = strings.ReplaceAll(str, "%", "")
	var u int = -1
	if strings.Contains(str, "-") {
		str = strings.ReplaceAll(str, "-", "")
	} else if strings.Contains(str, "+") {
		str = strings.ReplaceAll(str, "+", "")
		u = 1
	}
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Printf("%f of type %T", f, f)
	}

	return int(f) * u
}

func formatTickers(pair string, ticker Ticker) Tickers {
	volume := int(StringToFloat(ticker.Ticker.Volume))
	pricechangepercent := StringToInit(ticker.Ticker.PriceChangePercent)

	return Tickers{
		Pair:               pair,
		Volume:             volume,
		PriceChangePercent: pricechangepercent,
	}
}
func IntToString(n int) string {
	str := strconv.Itoa(n)
	return str
}
func PriceChangePercent(order_price string, market_price string) int {
	o := StringToFloat(order_price)
	m := StringToFloat(market_price)
	change := m/o*100 - 100
	return int(change)
}
