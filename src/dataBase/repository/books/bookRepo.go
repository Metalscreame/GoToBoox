package books

import (

	_ "github.com/lib/pq"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	"log"
	"fmt"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)



type BooksRepositoryPG struct{}

type BookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByCategory(categoryID int) ([]entity.Book, error)
	GetByID(bookID int) (entity.Book, error)
	GetMostPopularBooks(id int) ([]entity.Book, error)
}

//For connection to HerokuDatabase
/*func openDb() *sql.DB {
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
*/


func GetByID(bookID int) (entity.Book, error) {
	//for connection to HerokuDatabase
	//db := openDb()
	db:=dataBase.Connection

	rows, err := db.Query("SELECT title, description, popularity FROM gotoboox.books where id=$1", bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	book := entity.Book{}

	for rows.Next() {

		book = *new(entity.Book)
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
	return book, nil

}

func GetAll() ([]interface{}, error){
	//for connection to HerokuDatabase
	//db:=openDb()
		db:=dataBase.Connection

	rows, err := db.Query("SELECT a.title, a.description, a.popularity, b.title FROM gotoboox.books a, gotoboox.categories b where a.categoriesid=b.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	books := []interface{}{[]entity.Book{}, []entity.Categories{}}
	i:=0
	for rows.Next() {

		book := new(entity.Book)
		cat := new(entity.Categories)
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

func GetByCategory(categoryID int) ([]entity.Book, error) {
	//for connection to HerokuDatabase
	//db:=openDb()
	db:=dataBase.Connection
	rows, err := db.Query("SELECT title FROM gotoboox.books WHERE categoriesid=$1", categoryID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	books := []entity.Book{}
	i:=0

	for rows.Next() {

		book := new(entity.Book)
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
func (br BooksRepositoryPG) GetMostPopularBooks (id int) ([]entity.Book, error) {
	db := dataBase.Connection
	rows, err := db.Query("SELECT Id, Title, Popularity FROM gotoboox.books ORDER BY Popularity DESC LIMIT $ID", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var popularBooks []entity.Book
	for rows.Next() {
		var id int
		var title string
		var popularity float32
		err = rows.Scan(&id, &title, &popularity)
		if err != nil {
			return nil, err
		}
		book := entity.Book{ID: id, Title: title, Popularity: popularity}
		popularBooks = append(popularBooks, book)
	}
	return popularBooks, err
}




