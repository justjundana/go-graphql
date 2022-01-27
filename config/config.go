package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func FetchConnection() *sql.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/go-graphql?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
