package postgres

import (
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"database/sql"
	"errors"
)

type booksRepositoryPG struct {
	Db *sql.DB
}

func NewBooksRepository(Db *sql.DB) repository.BookRepository {
	return &booksRepositoryPG{Db}
}

const (
	takenState = "TAKEN"
)

//var Db = repository.OpenDb()
//GetByCategory iterates over the DB using the SQL SELECT Request and return selected book by its ID

func (p booksRepositoryPG) GetByID(bookID int) (books repository.Book, err error) {

	row := p.Db.QueryRow("SELECT title, description, popularity, state, image  FROM gotoboox.books where id = $1", bookID)
	if err != nil {
		log.Printf("Get %v", err)

		return
	}
	err = row.Scan(&books.Title, &books.Description, &books.Popularity, &books.State, &books.Image)

	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	//just for checking
	//fmt.Printf("%s\n%s\n%f\n", book.Title, book.Description, book.Popularity)
	return
}

func (p booksRepositoryPG) GetAllTakenBooks() (books []repository.Book, err error) {

	rows, err := p.Db.Query("SELECT id, title, description, image  FROM gotoboox.books where state = $1", takenState)
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	defer rows.Close()
	var book repository.Book
	for rows.Next() {

		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image);
			err != nil {
			log.Printf("Get %v", err)
		}
		books = append(books, book)
	}

	if err != nil {
		log.Printf("Get %v", err)
		return
	}

	return books, nil
}

//GetAll iterates over the DB using the SQL SELECT Request and return all books from DB
func (p booksRepositoryPG) GetAll() (books []repository.Book, err error) {
	rows, err := p.Db.Query("SELECT id, title, description, state  FROM gotoboox.books ")
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	defer rows.Close()
	var book repository.Book
	for rows.Next() {

		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.State);
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
	var book repository.Book
	for rows.Next() {

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
	rows, err := p.Db.Query("SELECT Id, Title, Popularity FROM books ORDER BY Popularity DESC LIMIT $1", quantity)
	if err != nil {
		return nil, errors.New("Failed to get the reply from a database")
	}
	defer rows.Close()
	var popularBooks []repository.Book
	var popBooks repository.Book
	for rows.Next() {
		err = rows.Scan(&popBooks.ID, &popBooks.Title, &popBooks.Popularity)
		if err != nil {
			return nil, errors.New("Failed to create the struct of books")
		}
		popularBooks = append(popularBooks, popBooks)
	}
	return popularBooks, nil
}

func (p booksRepositoryPG) InsertNewBook(b repository.Book) (err error) {
	_, err = p.Db.Query("INSERT INTO gotoboox.books (title,description,image) values($1,$2,$3)",
		b.Title, b.Description, b.Image)
	return
}

func (p booksRepositoryPG) UpdateBookState(bookId int, state string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.books set state=$1 where id=$2",
		state, bookId)
	return
}
