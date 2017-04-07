package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
	"errors"
	"database/sql"
)

type DBTest interface {
	Init(driver string) error
	Ping(timeout int) error
}

type DBInfo struct {
	Driver string
	DSN string
	Conn *sqlx.DB
}

func (db *DBInfo) Init(driver, DSN string) error {
	var err error
	db.Conn, err = sqlx.Open(driver, DSN)
	if err != nil {
		return errors.New("Error opening connection!")
	}
	return nil
}

func (db *DBInfo) Ping(timeout int) error {
	for try := 0;; try++ {
		if try > timeout {
			return errors.New("giving up")
		}
		if err := db.Conn.Ping(); err == nil {
			return nil
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var err error
	var result sql.Result

	db := DBInfo{}
	err = db.Init("postgres", "user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		fmt.Printf("db.Init failed: %s", err)
	}

	err = db.Ping(15)
	if err != nil {
		fmt.Printf("db.Ping failed: %s", err)
	} else {
		fmt.Println("Ping!")
	}

	schema := `CREATE TABLE place (
		country text,
		city text NULL,
		telcode integer);`

	result, err = db.Conn.Exec(schema)
	if err == nil {
		fmt.Printf("CREATE SCHEMA: %s", result)
	} else {
		fmt.Printf("ERROR :: CREATE SCHEMA: %s", err)
	}
}
