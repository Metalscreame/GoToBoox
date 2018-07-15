package dataBase

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"os"
)

//Connection is a global postgres connection variable, that must be used in all postgres repositories or DAO interfaces
var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
//For localhosts setup sys env "POSTGRES_URL" with key "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
//where ://username:password@host:port/dbname
func Connect() () {
	var err error
	dbURL, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		println("$POSTGRES_URL is required\nFor localhosts setup sys env \"POSTGRES_URL\" " +
			"with key \"postgres://postgres:root@localhost:5432/postgres?sslmode=disable\" where ://username:password@host:port/dbname")
		log.Fatal("$POSTGRES_URL is required\nFor localhosts setup sys env \"POSTGRES_URL\" " +
			"with key \"postgres://postgres:root@localhost:5432/postgres?sslmode=disable\" where ://username:password@host:port/dbname")
	}

	Connection, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	err = Connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}


//TokenKeyLookUp is a func that search for token value in system environment
func TokenKeyLookUp() (string) {
	tokenKey, ok := os.LookupEnv("TOKEN_KEY")
	if !ok {
		println("Missing tokenKey value\n Setup sys env \"TOKEN_KEY\" as any string you want and reload IDE")
		log.Fatal("Missing tokenKey value\n Setup sys env \"TOKEN_KEY\" as any string you want and reload IDE")
	}
	return tokenKey
}
