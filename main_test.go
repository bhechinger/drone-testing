package main

import (
	"testing"
	"fmt"
)

func TestPostgres(t *testing.T) {
	fmt.Println("Testing PostgreSQL",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	doit("postgres")
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing MySQL",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	doit("mysql")
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing Maria",
		"user=myapp host=localhost port=9876 password=mysecretpassword dbname=myapp sslmode=disable")
	doit("mysql")
}