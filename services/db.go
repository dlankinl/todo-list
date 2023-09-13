package services

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"reflect"
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

	fmt.Println("database was pinged successfully")
	rows, err := db.Query(`SELECT * FROM tasks`)
	for rows.Next() {
		var task Task

		s := reflect.ValueOf(&task).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)
		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		fmt.Println(1, task)
	}
}
