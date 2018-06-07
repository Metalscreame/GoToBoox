package postgres

import (
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"database/sql"
)

type booksRepositoryPG struct {
	Db *sql.DB
}

func NewBooksRepository(Db *sql.DB) repository.BookRepository {
	return &booksRepositoryPG{Db}
}

//var Db = repository.OpenDb()
//GetByCategory iterates over the DB using the SQL SELECT Request and return selected book by its ID
func (p booksRepositoryPG) GetByID(bookID int) (book repository.Book, err error) {

	rows := p.Db.QueryRow("SELECT title, description, popularity FROM gotoboox.books where id = $1", bookID)
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	if err != sql.ErrNoRows {
		log.Printf("Get %v", err)
		return
	}
	if err := rows.Scan(&book.Title, &book.Description, &book.Popularity);
		err != nil {
		log.Printf("Get %v", err)
	}
	//just for checking
	//fmt.Printf("%s\n%s\n%f\n", book.Title, book.Description, book.Popularity)
	//	fmt.Print(book)
	return book, err
}

//GetAll iterates over the DB using the SQL SELECT Request and return all books from DB
func (p booksRepositoryPG) GetAll() (books []repository.BookDescription, err error) {
	rows, err := p.Db.Query("SELECT a.title, a.description, a.popularity, b.title FROM gotoboox.books a, gotoboox.categories b where a.categories_id=b.id")
	if err != nil {
		log.Printf("Get %v", err)
	}
	defer rows.Close()
	var book repository.BookDescription
	for rows.Next() {

		if err := rows.Scan(&book.BookTitle, &book.Description, &book.Popularity, &book.CategoryTitle);
			err != nil {
			log.Printf("Get %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Get %v", err)
	}
	return books, nil

}

//GetByCategory iterates over the DB using the SQL SELECT Request and return books from chosen category
func (p booksRepositoryPG) GetByCategory(categoryID int) (books []repository.Book, err error) {
	rows, err := p.Db.Query("SELECT title FROM gotoboox.books WHERE categories_id=$1", categoryID)
	if err != nil {
		log.Printf("Get %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var book repository.Book
		if err := rows.Scan(&book.Title);
			err != nil {
			log.Printf("Get %v", err)
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Get %v", err)
	}
	return books, nil
}

//Function GetMostPopulareBooks iterates over the DB using the SQL SELECT Request and return 5 top-rated books.
func (p booksRepositoryPG) GetMostPopularBooks(quantity int) ([]repository.Book, error) {
	//	db := dataBase.Connection
	rows, err := p.Db.Query("SELECT Id, Title, Popularity FROM gotoboox.books ORDER BY Popularity DESC LIMIT $1", quantity)
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
