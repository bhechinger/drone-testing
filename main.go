package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
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

func doit(driver, DSN string) error {
	var err error
	var result sql.Result

	db := DBInfo{}
	err = db.Init(driver, DSN)
	if err != nil {
		return fmt.Errorf("db.Init failed: %s", err)
	}

	err = db.Ping(15)
	if err != nil {
		return fmt.Errorf("db.Ping failed: %s", err)
	}

	schema := `CREATE TABLE IF NOT EXISTS place (
		country text,
		city text NULL,
		telcode integer);`

	result, err = db.Conn.Exec(schema)
	if err != nil {
		return fmt.Errorf("CREATE SCHEMA ERROR: %s", err)
	}
	fmt.Printf("CREATE SCHEMA: %v\n", result)

	var query string
	query = `INSERT INTO place VALUES ('USA', 'Atlatla', 404)`
	result, err = db.Conn.Exec(query)
	if err != nil {
		return fmt.Errorf("INSERT ERROR: %s", err)
	}
	rows_a, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RowsAffected() ERROR: %s", err)
	}
	fmt.Printf("Rows inserted: %v\n", rows_a)

	type Place struct {
		Country       string
		City          sql.NullString
		TelephoneCode int `db:"telcode"`
	}

	p := Place{TelephoneCode: 404}
	stmt, err := db.Conn.PrepareNamed(`SELECT * FROM place WHERE telcode=:telcode`)
	rp := []Place{}
	err = stmt.Select(&rp, p)
	fmt.Printf("%+v\n", rp)

	return nil
}

func main() {
	if err := doit("mysql",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable"); err != nil {
		fmt.Printf("Something went wrong! %s", err)
	}
}
