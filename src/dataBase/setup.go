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
	DB_USER     string `json:"db_user"`
	DB_PASSWORD string `json:"db_password"`
	DB_NAME     string `json:"db_name"`
	DB_HOST     string `json:"db_host"`
	DB_PORT     string `json:"db_port"`
}

var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
func Connect(d DataBaseCredentials) () {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", d.DB_USER, d.DB_PASSWORD, d.DB_HOST, d.DB_PORT, d.DB_NAME)
	Connection, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	Connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}
