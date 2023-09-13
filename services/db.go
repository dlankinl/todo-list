package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

var db *sql.DB

const dbTimeout = time.Second * 3

func InitConnection(connStr string) *DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while connecting to db:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("error while pinging db:", err)
	}

	fmt.Println("database was pinged successfully", db)
	dbConn.DB = db
	return dbConn
}

func New(dbPool *sql.DB) {
	db = dbPool
}
