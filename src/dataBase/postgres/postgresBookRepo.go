package postgres

import (
	"log"
	"fmt"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"database/sql"
)

type booksRepositoryPG struct{
	Db *sql.DB
}

func NewBooksRepository(Db *sql.DB) repository.BookRepository {
	return &booksRepositoryPG{Db}
}

//GetByCategory iterates over the DB using the SQL SELECT Request and return selected book by its ID
func (p booksRepositoryPG) GetByID(bookID int) (b repository.Book, err error) {
	//for connection to HerokuDatabase
	//db := repository.OpenDb()
//	db:=dataBase.Connection

	rows, err := p.Db.Query("SELECT title, description, popularity FROM gotoboox.books where id=$1", bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	book := repository.Book{}

	for rows.Next() {

		book = *new(repository.Book)
		if err := rows.Scan(&book.Title, &book.Description, &book.Popularity);
			err != nil {
			log.Fatal(err)
		}

		//just for checking
		fmt.Printf("%s\n%s\n%f\n", book.Title, book.Description, book.Popularity)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return

}


//GetAll iterates over the DB using the SQL SELECT Request and return all books from DB
func (p booksRepositoryPG) GetAll() ([]interface{}, error) {
	//for connection to HerokuDatabase
	//db:=openDb()
//	db:=dataBase.Connection
	//db := repository.OpenDb()
	rows, err := p.Db.Query("SELECT a.title, a.description, a.popularity, b.title FROM gotoboox.books a, gotoboox.categories b where a.categoriesid=b.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	books := []interface{}{[]repository.Book{}, []repository.Categories{}}
	i:=0
	for rows.Next() {

		book := new(repository.Book)
		cat := new(repository.Categories)
		if err := rows.Scan(&book.Title, &book.Description, &book.Popularity, &cat.Title );
			err != nil {
			log.Fatal(err)
		}
		books = append(books, *book, *cat)
		i++
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil

}

//GetByCategory iterates over the DB using the SQL SELECT Request and return books from chosen category
func (p booksRepositoryPG)  GetByCategory(categoryID int) ([]repository.Book, error) {
	//for connection to HerokuDatabase
	//db := repository.OpenDb()
//	db:=dataBase.Connection
	rows, err := p.Db.Query("SELECT title FROM gotoboox.books WHERE categoriesid=$1", categoryID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []repository.Book{}
	i:=0

	for rows.Next() {

		book := new(repository.Book)
		if err := rows.Scan(&book.Title);
			err != nil {
			log.Fatal(err)
		}
		books = append(books, *book)
		i++
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil
}

//Function GetMostPopulareBooks iterates over the DB using the SQL SELECT Request and return 5 top-rated books.
func (p booksRepositoryPG) GetMostPopularBooks (id int) ([]repository.Book, error) {
//	db := dataBase.Connection
	rows, err := p.Db.Query("SELECT Id, Title, Popularity FROM gotoboox.books ORDER BY Popularity DESC LIMIT $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var popularBooks []repository.Book
	for rows.Next() {
		var id int
		var title string
		var popularity float32
		err = rows.Scan(&id, &title, &popularity)
		if err != nil {
			return nil, err
		}
		book := repository.Book{ID: id, Title: title, Popularity: popularity}
		popularBooks = append(popularBooks, book)
	}
	return popularBooks, err
}
