package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	connStr := "user=postgres password=psql dbname=product sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(db)
	} else {
		log.Println("ошибка в db")
	}
}
