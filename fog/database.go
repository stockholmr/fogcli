package fog

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stockholmr/logging"
)

var _db *sqlx.DB

func Connect(hostname string, username string, password string, database string) {
	var err error

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username,
		password,
		hostname,
		3306,
		database,
	)

	_db, err = sqlx.Connect("mysql", connectionString)
	if err != nil {
		logging.Fatal(err.Error())
	}
}

func Disconnect() {
	_db.Close()
}
