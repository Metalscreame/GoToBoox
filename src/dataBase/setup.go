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
}

var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
func Connect(d DataBaseCredentials) (*sql.DB) {
	var err error
	//temporary heroku solution
	d.DB_USER="zrlfyamblttpom"
	d.DB_PASSWORD="e2c0e8832ea228e6b15e553ce69f7cb2c0ff4d646ff0f284245ce77cc78b437b"
	d.DB_NAME = "d7ckgvm53enhum"
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
