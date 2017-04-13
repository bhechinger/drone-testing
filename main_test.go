package main

import (
	"testing"
)

func TestPostgres(t *testing.T) {
	doit("postgres")
}

func TestMysql(t *testing.T) {
	doit("mysql")
}