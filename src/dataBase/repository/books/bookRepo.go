package books

import (

	_ "github.com/lib/pq"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	"database/sql"
	"log"
	"fmt"
)

/*type Book struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Popularity int `json:"popularity"`
	CategoriesID int `json:"categoriesID"`
}*/




type BookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByCategory(categoryID int) ([]entity.Book, error)
	GetMostPopular() ([]entity.Book, error)
}


func openDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://zrlfyamblttpom:e2c0e8832ea228e6b15e553ce69f7cb2c0ff4d646ff0f284245ce77cc78b437b@ec2-54-247-111-19.eu-west-1.compute.amazonaws.com:5432/d7ckgvm53enhum")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("Can not connect to database: ", err)
	}

	return db
}

func GetAll() ([]entity.Book, error){
	db:=openDb()

	rows, err := db.Query("SELECT title FROM gotoboox.books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []entity.Book{}
	i:=0
//	fmt.Printf("Titles: \n")
	for rows.Next() {

		book := new(entity.Book)
		if err := rows.Scan(&book.Title);
		err != nil {
			log.Fatal(err)
		}
		books = append(books, *book)
		fmt.Printf("%s\n", books[i].Title)
		i++
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil



}



func GetByCategory(categoryID int) ([]entity.Book, error) {
	db:=openDb()
	rows, err := db.Query("SELECT title FROM gotoboox.books WHERE categoriesid=$1", categoryID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []entity.Book{}
	i:=0
	//	fmt.Printf("Titles: \n")
	for rows.Next() {

		book := new(entity.Book)
		if err := rows.Scan(&book.Title);
			err != nil {
			log.Fatal(err)
		}
		books = append(books, *book)
		fmt.Printf("%s\n", books[i].Title)
		i++
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil
}

func GetByPopularity() ([]entity.Book, error) {

	db := openDb()
	rows, err := db.Query("SELECT title, popularity  FROM gotoboox.books ORDER BY popularity DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []entity.Book{}
	i := 0
	//	fmt.Printf("Titles - popularity \n")
	for rows.Next() {

		book := new(entity.Book)
		if err := rows.Scan(&book.Title, &book.Popularity);
			err != nil {
			log.Fatal(err)
		}
		books = append(books, *book)
		fmt.Printf("%s - %d booXmark\n", books[i].Title, books[i].Popularity)

	i++

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil

}





