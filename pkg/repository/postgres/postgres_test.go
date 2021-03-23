package postgres_test

import (
	"testing"

	"github.com/nicholaslim94/messenger_backend/pkg/repository/postgres"
)

func TestConnect(t *testing.T) {
	db, err := postgres.Connect("localhost", 5432, "messenger", "postgres", "password", false)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	t.Log("Db connection established ok")
}
