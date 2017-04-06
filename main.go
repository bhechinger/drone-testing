package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	doit()
}

func doit() {
	db, err := sqlx.Open("postgres",
		"user=myapp host=localhost dbname=myapp sslmode=disable")
	if err != nil {
		fmt.Println("Error Opening!")
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging db: %s\n", err)
	} else {
		fmt.Println("It works!")
	}
}
