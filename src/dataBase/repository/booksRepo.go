package repository

import (
	"fmt"
	"database/sql"
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/configuration"
	"github.com/metalscreame/GoToBoox/src/dataBase/entity"
)

type Books entity.Book

func GetBooks() []Books{
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		configuration.DB_USER, configuration.DB_PASSWORD, configuration.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	books := []Books{}
	for rows.Next() {
		book := new(Books)
		if err := rows.Scan(&book.Id, &book.Title); err != nil {
			log.Fatal(err)
		}
		books = append(books, *book)
		fmt.Printf("id %d title %s\n", books[0].Id, books[0].Title)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books
}

