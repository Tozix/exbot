package mysql

import (
	"database/sql"
	"exbot/exb"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func NewSql(User string, Pass string, Host string, Name string) (*Sql, error) {
	db, err := sql.Open("mysql", User+":"+Pass+"@tcp("+Host+":3306)/"+Name)
	checkError(err)
	//defer db.Close()

	err = db.Ping()
	checkError(err)
	log.Println("Successfully created connection to database.")
	return &Sql{
		DB: db,
	}, nil
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Insert - Метод для создания новой заметки в базе дынных.
func (s *Sql) InsertOrder(order exb.Order, table string) (int, error) {
	// Ниже будет SQL запрос, который мы хотим выполнить. Мы разделили его на две строки
	// для удобства чтения (поэтому он окружен обратными кавычками
	// вместо обычных двойных кавычек).
	stmt := "insert ignore into " + table + "(ID,UUID,Side,OrdType,Price,AvgPrice,State,Market,MarketType,CreatedAt,UpdatedAt,OriginVolume) values (?,?,?,?,?,?,?,?,?,?,?,?)"

	// Используем метод Exec() из встроенного пула подключений для выполнения
	// запроса. Первый параметр это сам SQL запрос, за которым следует
	// заголовок заметки, содержимое и срока жизни заметки. Этот
	// метод возвращает объект sql.Result, который содержит некоторые основные
	// данные о том, что произошло после выполнении запроса.
	result, err := s.DB.Exec(stmt,
		order.ID,
		order.UUID,
		order.Side,
		order.OrdType,
		order.Price,
		order.AvgPrice,
		order.State,
		order.Market,
		order.MarketType,
		order.CreatedAt,
		order.UpdatedAt,
		order.OriginVolume)
	if err != nil {
		return 0, err
	}

	// Используем метод LastInsertId(), чтобы получить последний ID
	// созданной записи из таблицу snippets.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Возвращаемый ID имеет тип int64, поэтому мы конвертируем его в тип int
	// перед возвратом из метода.
	return int(id), nil
}

func (s *Sql) OrdersList(table string) (*sql.Rows, error) {
	rows, err := s.DB.Query("select * from " + table)
	if err != nil {
		panic(err)
	}
	return rows, nil

}
