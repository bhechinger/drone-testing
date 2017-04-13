package main

import (
	"testing"
	"fmt"
)

func TestPostgres(t *testing.T) {
	fmt.Println("Testing PostgreSQL")
	err := doit("postgres",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		t.Errorf("PostgreSQL failed: %v", err)
	}
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing MySQL")
	err := doit("mysql",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		t.Errorf("MySQL failed: %v", err)
	}
}

func TestMariadb(t *testing.T) {
	fmt.Println("Testing Maria")
	err := doit("mysql",
		"user=myapp host=localhost port=9876 password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		t.Errorf("MariaDB failed: %v", err)
	}
}