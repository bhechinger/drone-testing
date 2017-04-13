package main

import (
	"testing"
	"fmt"
)

func TestPostgres(t *testing.T) {
	fmt.Println("Testing PostgreSQL")
	doit("postgres",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing MySQL")
	doit("mysql",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing Maria")
	doit("mysql",
		"user=myapp host=localhost port=9876 password=mysecretpassword dbname=myapp sslmode=disable")
}