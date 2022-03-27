package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"lead_generation_basic/config"
)

var Db = connect()

// connect setup postgres sql client
func connect() *sqlx.DB {
	pSqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetString("sqldb.host"),
		config.GetInt("sqldb.port"),
		config.GetString("sqldb.user"),
		config.GetString("sqldb.password"),
		config.GetString("sqldb.dbname"))

	db, err := sqlx.Connect("postgres", pSqlConn)
	if err != nil {
		panic("Failed to initialize db client: " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("Failed to ping db: " + err.Error())
	}

	return db
}
