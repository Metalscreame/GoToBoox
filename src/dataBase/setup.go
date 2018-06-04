package dataBase

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	DB_SCHEMA              = "gotoboox."
	DB_USERS_TABLE         = "users"
	DB_BOOKS_TABLE         = "books"
	DB_AUTHORS_TABLE       = "authors"
	DB_CATEGORIES_TABLE    = "categories"
	DB_BOOKS_AUTHORS_TABLE = "books_authors"
)

type DataBaseCredentials struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
func Connect(d DataBaseCredentials) (*sql.DB) {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", d.DB_USER, d.DB_PASSWORD, d.DB_NAME)
	Connection, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	Connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return Connection
}
