package dataBase

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"os"
)



var Connection *sql.DB

//Connect is a function that is used to open Connection
//with a dataBase.
//For localhosts setup sys env "POSTGRES_URL" with key "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
//where ://username:password@host:port/dbname
func Connect() () {
	var err error
	dbUrl, ok := os.LookupEnv("POSTGRES_URL")
	if !ok {
		println("$POSTGRES_URL is required\nFor localhosts setup sys env \"POSTGRES_URL\" " +
			"with key \"postgres://postgres:root@localhost:5432/postgres?sslmode=disable\" where ://username:password@host:port/dbname")
		log.Fatal("$POSTGRES_URL is required\nFor localhosts setup sys env \"POSTGRES_URL\" " +
			"with key \"postgres://postgres:root@localhost:5432/postgres?sslmode=disable\" where ://username:password@host:port/dbname")

	}

	Connection, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	err=Connection.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return
}
