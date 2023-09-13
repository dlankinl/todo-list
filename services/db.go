package services

import (
	"database/sql"
	"log"
	"time"
)

var db *sql.DB

const dbTimeout = time.Second * 3

func InitConnection(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while connecting to db:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("error while pinging db:", err)
	}
}
