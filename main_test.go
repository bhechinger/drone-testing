package main

import (
	"testing"
	"fmt"
)

func TestPostgres(t *testing.T) {
	fmt.Println("Testing PostgreSQL")
	doit("postgres")
}

func TestMysql(t *testing.T) {
	fmt.Println("Testing MySQL")
	doit("mysql")
}