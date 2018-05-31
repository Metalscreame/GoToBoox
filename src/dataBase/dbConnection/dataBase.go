package dbConnection

import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pq"
)

//Specify this values for concrette db
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "root"
	DB_NAME     = "postgres"
	DB_SCHEMA   = "gotoboox."
	DB_USERS_TABLE = "users"
	DB_BOOKS_TABLE = "books"
	DB_AUTHORS_TABLE = "authors"
	DB_CATEGORIES_TABLE = "categories"
	DB_BOOKS_AUTHORS_TABLE = "books_authors"
)

//GlobalDataBaseConnection is a global variableto manage connections to database
var GlobalDataBaseConnection *sql.DB

//InitializeConnection is a function that is used to open connection
//with a dataBase.
func InitializeConnection()  {
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	GlobalDataBaseConnection, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
}
