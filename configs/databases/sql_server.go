package databases

import (
	"database/sql"
	"esaku-project/configs"
	"esaku-project/helpers"
	"fmt"
	"time"
)

func NewSqlServer(configuration configs.Config) *sql.DB {
	var (
		database = configuration.Get("DB_SQL_SERVER_DATABASE")
		port     = configuration.Get("DB_SQL_SERVER_PORT")
		host     = configuration.Get("DB_SQL_SERVER_HOST")
		username = configuration.Get("DB_SQL_SERVER_USERNAME")
		password = configuration.Get("DB_SQL_SERVER_PASSWORD")
	)

	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", username, password, host, port, database)
	db, err := sql.Open("sqlserver", conn)

	helpers.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
