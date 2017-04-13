package main

import (
	"testing"
	"fmt"
)

func TestDB(t *testing.T) {
	fmt.Println("Testing DB")
	err := doit("postgres",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		t.Errorf("DB Test failed: %v", err)
	}
}
