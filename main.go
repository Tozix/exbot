package main

import (
	"exbot/exb"
	"log"
)

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
	//balance := trade.Balance("avn").Balance
	//log.Println("Баланс монеты:" + balance)

	//trades := trade.SellOrder("avnusdt", 11000.0, 0.0011)
	//cancel := trade.CancelOrders("avnusdt")
	orders, err := trade.GetOpenOrders("avnusdt")
	if err != nil {
		log.Fatalln("Какая-то херня: " + err.Error())
	}

	for _, order := range orders {
		if order.State == "wait" {
			log.Printf("IDы: %v Магаз: %v Объеб монет: %v", order.ID, order.Market, order.OriginVolume)
		}
	}

	/*
		adr := trade.GetAddress(config.Dex.Quote_asset).Available
		log.Println(adr)
	*/
}
