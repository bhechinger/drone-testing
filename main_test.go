package main

import (
	"testing"
	"fmt"
	"os"
)

func TestDB(t *testing.T) {
	fmt.Println("Testing DB")
	fmt.Println("DATABASE:", os.Getenv("DATABASE"))
	fmt.Println("GO_VERSION:", os.Getenv("GO_VERSION"))
	err := doit("postgres",
		"user=myapp host=localhost password=mysecretpassword dbname=myapp sslmode=disable")
	if err != nil {
		t.Errorf("DB Test failed: %v", err)
	}
}
