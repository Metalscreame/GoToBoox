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


type BooksRepository struct{

}

type BookRepository interface {
	GetAll() ([]entity.Book, error)
	GetByCategory(categoryID int) ([]entity.Book, error)
	GetMostPopular() ([]entity.Book, error)
	GetByID(bookID int) (entity.Book, error)
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



func GetByID(bookID int) (entity.Book, error) {
	db:=openDb()

	rows, err := db.Query("SELECT title, description, popularity FROM gotoboox.books where id=$1", bookID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	book := entity.Book{}

	for rows.Next() {

		book := new(entity.Book)
		if err := rows.Scan(&book.Title);
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
	db:=openDb()

	rows, err := db.Query("SELECT a.title, a.description, a.popularity, b.title FROM gotoboox.books a, gotoboox.categories b where a.categoriesid=b.id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	books := []interface{}{[]entity.Book{}, []entity.Categories{}}
	//books, cats := []entity.Book{}, []entity.Categories{}
	i:=0
//	fmt.Printf("Titles: \n")
	for rows.Next() {

		book := new(entity.Book)
		cat := new(entity.Categories)
		if err := rows.Scan(&book.Title, &book.Description, &book.Popularity, &cat.Title );
		err != nil {
			log.Fatal(err)
		}
		books = append(books, *book, *cat)

		//just for checking
		fmt.Printf("%v\n%v\n", book.Title, cat.Title)
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
		//just for checking
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
		//just for checking
		fmt.Printf("%s - %d booXmark\n", books[i].Title, books[i].Popularity)

	i++

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return books, nil

}





