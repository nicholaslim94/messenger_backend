package postgres

import (
	"database/sql"
	"fmt"
	"log"

	//Postgres driver
	_ "github.com/lib/pq"
)

//Connect opens a connection with postgres and ping it to check its connection status
func Connect(host string, port int, dbName string, user string,
	password string, sslMode bool) (*sql.DB, error) {
	ssl := "disable"
	if sslMode {
		ssl = "enable"
	}
	dbDetails := fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=%s&sslmode=%s",
		host, port, dbName, user, password, ssl)
	dbLog := fmt.Sprintf("postgresql://%s:%d/%s?user=%s&password=*&sslmode=%s",
		host, port, dbName, user, ssl)
	log.Println(dbLog)
	db, err := sql.Open("postgres", dbDetails)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
