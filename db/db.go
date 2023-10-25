package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

//type Conn struct {
//	db *sql.DB
//}

func ConnectDB(connStr string) *sql.DB {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error while connecting to db:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("error while pinging db:", err)
	}

	fmt.Println("database was pinged successfully", db)
	return db
}
