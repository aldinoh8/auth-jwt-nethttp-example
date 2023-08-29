package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDb(dbcreds string) *sql.DB {
	db, err := sql.Open("mysql", dbcreds)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
