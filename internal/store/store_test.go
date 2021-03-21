package store_test

import (
	"os"
	"testing"
)

var databaseURL string

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		databaseURL = "host=localhost dbname=e-shop sslmode=disable user=bakhodur password=1996"
	}

	os.Exit(m.Run())
}
