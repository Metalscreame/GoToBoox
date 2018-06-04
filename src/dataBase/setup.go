package dataBase

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

//Specify this values for concrette db
const (
	//DB_USER     ="zrlfyamblttpom"//"postgres"//  // for local
	//DB_PASSWORD = "e2c0e8832ea228e6b15e553ce69f7cb2c0ff4d646ff0f284245ce77cc78b437b"//"root"//
	//DB_NAME     = "d7ckgvm53enhum" //"postgres"//
	//
	////use this for local machines
	////DB_USER                = "postgres"
	////DB_PASSWORD            = "root"
	////DB_NAME                = "postgres"
	DB_SCHEMA              = "gotoboox."
	DB_USERS_TABLE         = "users"
	DB_BOOKS_TABLE         = "books"
	DB_AUTHORS_TABLE       = "authors"
	DB_CATEGORIES_TABLE    = "categories"
	DB_BOOKS_AUTHORS_TABLE = "books_authors"
)

type DataCredentials struct {
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
}
var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
func Connect(dbConf DataCredentials) (*sql.DB){
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbConf.DB_USER, dbConf.DB_PASSWORD, dbConf.DB_NAME)
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
