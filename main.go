package main

import (
	"exbot/exb"
	"log"
)

/* АЛгоритм */
/*
1) Проверяем открытые ордеры из базы (Ордера в базе нет, но он открыт - заносим в базу)
2) Дрочим базу постояно не исполнен ли ордер и мониторим изменение цены (Возможно монета сильно просела и ее стоит усреднить)
Если по нужной паре нет открытых ордеров инцииализируем открытие нового ордера.
Пинаем вайт лист, подбераем кандидатов на новый ордер удовлетворяющий требованиям
Закупаем его по текущей цене (На минимальный объем указанный в конфиге) - Выставляем ордер (заносим в базу)




*/

func main() {

	config, err := GetConfig("config.json")
	if err != nil {
		log.Fatalln("Невозможно загрузить файл конфигурации:", err.Error())
	}
	//Ключи для начала торговли
	trade, err := exb.NewTrade(config.Exb.Key, config.Exb.Secret)
	if err != nil {
		log.Fatalln("Какая-то херня: " + err.Error())
	}
	//Доступный баланс
	balance := trade.Balance("avn").Balance
	log.Println("Баланс монеты:" + balance)

	//trades := trade.SellOrder("avnusdt", 11000.0, 0.0011)
	//cancel := trade.CancelOrders("avnusdt")

	orders, err := trade.GetOpenOrders("avnusdt")
	if err != nil {
		log.Fatalln("Какая-то херня: " + err.Error())
	}

	for _, order := range orders {
		//if order.State == "wait" {
		log.Printf("IDы: %v Магаз: %v Объеб монет: %v", order.ID, order.Market, order.OriginVolume)
		//	}
	}

	//log.Println(tickers.PriceChangePercent)
	tickers := trade.GetAllTickers(config.Exb.Quote_asset)
	for _, ticker := range tickers {
		if ticker.Volume > config.Exb.Min_volume {
			log.Printf("Пара: %s Объем торгов: %d Изменение цены: %d", ticker.Pair, ticker.Volume, ticker.PriceChangePercent)
		}
	}
	//log.Println(tickers)

	/*
		adr := trade.GetAddress(config.Dex.Quote_asset).Available
		log.Println(adr)
	*/
}
