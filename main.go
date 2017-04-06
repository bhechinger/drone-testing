package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

func main() {
	doit()
}

func doit() {
	db, err := sqlx.Open("postgres",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		fmt.Println("Error Opening!")
	}

	for try := 0;; try++ {
		fmt.Println(try)
		if try > 100 {
			fmt.Println("giving up")
			break
		}
		if err := db.Ping(); err != nil {
			fmt.Printf("Error pinging db: %s\n", err)
		} else {
			fmt.Println("It works!")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
