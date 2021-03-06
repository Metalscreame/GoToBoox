package postgres

import (
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"database/sql"
	"errors"
	"strings"
	"github.com/lib/pq"
)

type booksRepositoryPG struct {
	Db *sql.DB
}
//NewBooksRepository is a function to get New BooksRepository which uses given connection
func NewBooksRepository(Db *sql.DB) repository.BookRepository {
	return &booksRepositoryPG{Db}
}

const (
	takenState = "TAKEN"
)

//InsertTags inserts to database tags that belong to book
func (p *booksRepositoryPG) InsertTags(tagID int, bookID int) (err error) {

	_, err = p.Db.Query("INSERT INTO gotoboox.books_tags (id, tag_id) values($1, $2)", bookID, tagID)
	return
}

//GetTagsForBook get list of tags that belong to certain book
func (p *booksRepositoryPG) GetTagsForBook(bookID int) (tags []repository.Book, err error) {

	rows, err := p.Db.Query("SELECT gotoboox.tags.title FROM (gotoboox.books JOIN gotoboox.books_tags USING (id)) JOIN gotoboox.tags USING (tag_id) WHERE gotoboox.books.id = $1", bookID)
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	defer rows.Close()
	var tag repository.Book
	for rows.Next() {

		if err := rows.Scan(&tag.TagsTitles);
			err != nil {
			log.Printf("Get %v", err)
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Get %v", err)

	}
	return tags, nil
}

//GetByID iterates over the DB using the SQL SELECT Request and return selected book by its ID
func (p *booksRepositoryPG) GetByID(bookID int) (books repository.Book, err error) {

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
	return
}

//GetAllTakenBooks iterates over the DB using the SQL SELECT Request and return all books with state "taken"
func (p *booksRepositoryPG) GetAllTakenBooks() (books []repository.Book, err error) {

	rows, err := p.Db.Query("SELECT id, title, description, image  FROM gotoboox.books where state = $1", takenState)
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	defer rows.Close()
	var book repository.Book
	for rows.Next() {
		if err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image);
			err != nil {
			log.Printf("Get %v", err)
		}
		books = append(books, book)
	}

	if err != nil {
		log.Printf("Got %v", err)
		return
	}

	return books, nil
}

//GetAll iterates over the DB using the SQL SELECT Request and return all books from DB
func (p *booksRepositoryPG) GetAll() (books []repository.Book, err error) {
	rows, err := p.Db.Query("SELECT id, title, description, state  FROM gotoboox.books LIMIT 100")
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
func (p *booksRepositoryPG) GetByCategory(categoryID int) (books []repository.Book, err error) {
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

//GetByLikeName iterates over the DB using the SQL SELECT Request and return books by name
func (p booksRepositoryPG) GetByLikeName(title string) (books []repository.Book, err error) {
	rows, err := p.Db.Query("SELECT id, title FROM gotoboox.books WHERE LOWER(title) LIKE '%' || $1 || '%'", strings.ToLower(title))
	if err != nil {
		log.Printf("Get %v", err)
	}
	defer rows.Close()
	var book repository.Book
	for rows.Next() {

		if err := rows.Scan(&book.ID, &book.Title);
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

func getResult(rows *sql.Rows, rowsError error) (books []repository.Book, err error) {
	defer rows.Close()
	if rowsError != nil {
		log.Printf("Get %v", err)
	}

	var book repository.Book
	for rows.Next() {
		books = append(books, book)
		if err := rows.Err(); err != nil {
			log.Printf("Get %v", err)
		}
	}
	return books, err
}

//GetByTagsAndRating iterates over the DB using the SQL SELECT Request and return books by tags AND/OR rating
func (p booksRepositoryPG) GetByTagsAndRating(tags []string, rating []int) (books []repository.Book, err error) {
	tagsLen := len(tags)
	// if user don't select rating, but select the tags
	if rating[0] >= 0 && rating[1] >= 0 {
		if rating[0] == 0 && rating[1] == 0 && tagsLen != 0 {
			rows, err := p.Db.Query("SELECT gotoboox.books.id, gotoboox.books.title FROM gotoboox.books "+
				"LEFT JOIN gotoboox.books_tags ON gotoboox.books.id = gotoboox.books_tags.id "+
				"LEFT JOIN gotoboox.tags ON gotoboox.books_tags.tag_id = gotoboox.tags.tag_id "+
				"WHERE gotoboox.tags.title = any($1) "+
				"GROUP BY gotoboox.books.title, gotoboox.books.id "+
				"having count(*) = $2",
				pq.Array(tags), tagsLen)
			return getResult(rows, err)
	} else if tagsLen == 0 && rating[0] != 0 && rating[1] != 0 {
		// if user don't select tags, but select the rating
		rows, err := p.Db.Query("SELECT gotoboox.books.id, gotoboox.books.title FROM gotoboox.books "+
			"LEFT JOIN gotoboox.books_tags ON gotoboox.books.id = gotoboox.books_tags.id "+
			"LEFT JOIN gotoboox.tags ON gotoboox.books_tags.tag_id = gotoboox.tags.tag_id "+
			"WHERE gotoboox.books.popularity BETWEEN $1 AND $2"+
			"GROUP BY gotoboox.books.title, books.id ",
			rating[0], rating[1])
			return getResult(rows, err)
	} else {
		// if user select the rating and tags
		rows, err := p.Db.Query("SELECT gotoboox.books.id, gotoboox.books.title FROM gotoboox.books "+
			"LEFT JOIN gotoboox.books_tags ON gotoboox.books.id = gotoboox.books_tags.id "+
			"LEFT JOIN gotoboox.tags ON gotoboox.books_tags.tag_id = gotoboox.tags.tag_id "+
			"WHERE gotoboox.tags.title = any($1) AND gotoboox.books.popularity BETWEEN $3 AND $4"+
			"GROUP BY gotoboox.books.title, gotoboox.books.id "+
			"having count(*) = $2",
			pq.Array(tags), tagsLen, rating[0], rating[1])
			return getResult(rows, err)
		}
	}
	return books, nil
}

//Function GetMostPopulareBooks iterates over the DB using the SQL SELECT Request and return 5 top-rated books.
func (p booksRepositoryPG) GetMostPopularBooks(quantity int) ([]repository.Book, error) {
	rows, err := p.Db.Query("SELECT Id, Title, Popularity FROM gotoboox.books ORDER BY Popularity DESC LIMIT $1", quantity)
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

func (p *booksRepositoryPG) InsertNewBook(b repository.Book) (lastID int, err error) {
	err = p.Db.QueryRow("INSERT INTO gotoboox.books (title,description,image, popularity) values($1,$2,$3, $4) RETURNING id",
		b.Title, b.Description, b.Image, b.Popularity).Scan(&lastID)
	return
}

func (p *booksRepositoryPG) UpdateBookState(bookID int, state string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.books set state=$1 where id=$2",
		state, bookID)
	return
}

func (p *booksRepositoryPG) UpdateBookStateAndUsersBookIDByUserEmail(email string, state string, bookID int) (err error) {
	tx, err := p.Db.Begin()
	if err != nil {
		return
	}

	//First transaction
	{
		stmt, err := tx.Prepare(`UPDATE gotoboox.users set book_id=$1, has_book_for_exchange=TRUE where email=$2`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(bookID, email); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return err
		}
	}
	//Second transaction
	{
		stmt, err := tx.Prepare(`UPDATE gotoboox.books set  state=$1 where id=$2`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(state, bookID); err != nil {
			tx.Rollback() // return an error too, we may want to wrap them
			return err
		}
	}
	return tx.Commit()
}
