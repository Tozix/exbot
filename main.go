package main

import (
	"exbot/exb"
	"exbot/mysql"
	"log"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := mysql.NewSql(config.MySQL.User, config.MySQL.Pass, config.MySQL.Host, config.MySQL.Name)
	if err != nil {
		log.Fatalln("Какая-то херня: " + err.Error())
	}
	defer db.DB.Close()

	//Доступный баланс
	quote_asset_balance := trade.Balance(config.Exb.Quote_asset).Balance
	log.Println("Доступный баланс для закупа:" + quote_asset_balance)

	//trades := trade.SellOrder("avnusdt", 11000.0, 0.0011)
	//cancel := trade.CancelOrders("avnusdt")

	orders, err := trade.GetOpenOrders("avnusdt")
	if err != nil {
		log.Fatalln("Какая-то херня: " + err.Error())
	}

	//Заносим в нашу базу
	for _, order := range orders {
		log.Printf("IDы: %v", order)
		_, err := db.InsertOrder(order, "orders")
		if err != nil {
			panic(err)
		}
	}

	rows, err := db.OrdersList("orders")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var db_order exb.Order
		err = rows.Scan(&db_order.ID,
			&db_order.UUID,
			&db_order.Side,
			&db_order.OrdType,
			&db_order.Price,
			&db_order.AvgPrice,
			&db_order.State,
			&db_order.Market,
			&db_order.MarketType,
			&db_order.CreatedAt,
			&db_order.UpdatedAt,
			&db_order.OriginVolume,
			&db_order.RemainingVolume,
			&db_order.ExecutedVolume,
			&db_order.MakerFee,
			&db_order.TakerFee,
			&db_order.TradesCount)
		if err != nil {
			panic(err.Error())
		}
		//log.Printf(db_order.UUID)
		order := trade.GetOrder(exb.IntToString(db_order.ID))
		if db_order.State == "wait" && db_order.Side == "sell" {
			//Если статус ждет, запрашиваем текущую цену по этой моменте
			ticker := trade.GetTicker(order.Market).Ticker
			log.Printf("Прайс в ордере: %v", order.Price)
			log.Printf("Last: %s", ticker.Last)
			log.Printf("Изменение цены: %d", exb.PriceChangePercent(order.Price, ticker.Last))
		}

	}

	//log.Println(tickers.PriceChangePercent)
	//tickers := trade.GetAllTickers(config.Exb.Quote_asset)
	//for _, ticker := range tickers {
	//	if ticker.Volume > config.Exb.Min_volume {
	//		log.Printf("Пара: %s Объем торгов: %d Изменение цены: %d", ticker.Pair, ticker.Volume, ticker.PriceChangePercent)
	//	}
	//}

}
