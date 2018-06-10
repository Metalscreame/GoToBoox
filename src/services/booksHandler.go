package services

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
	"encoding/base64"
	"strings"
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

type BookService struct {
	BooksRepo repository.BookRepository
}

func NewBookService(repository repository.BookRepository) *BookService {
	return &BookService{
		BooksRepo: repository,
	}
}

func (b *BookService) FiveMostPop(c *gin.Context) {
	FiveMostPop, _ := b.BooksRepo.GetMostPopularBooks(5)
	if len(FiveMostPop) > 0 {
		c.JSON(http.StatusOK, FiveMostPop)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books have been found"})
	}
}

//showAllBooks is a handler for GetAll function
func (b *BookService) showAllBooks(c *gin.Context) {
	type Data struct{

		Book []repository.Book
	}
	books, err := b.BooksRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		output := Data{books}
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
		return
	}
}

//getBooks is a handler for GetByCategory function
func (b *BookService) getBooks(c *gin.Context) {
	// Check if the categoryID is valid
	if catID, err := strconv.Atoi(c.Param("cat_id"));
		err != nil {
		// If the book is not found, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		// Check if the category exists
		if book, err := b.BooksRepo.GetByCategory(catID);
			err != nil {
			// If an invalid category ID is specified in the URL, abort with an error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, book)
			return
		}
	}
}

//getBook is a handler for GetByID function
func (b *BookService) getBook(c *gin.Context) {
	type Data struct{

		Book repository.Book
	}
	// Check if the bookID is valid
	if bookID, err := strconv.Atoi(c.Param("book_id"));

		err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {

		// Check if the category exists

		if book, err := b.BooksRepo.GetByID(bookID);
			err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			output := Data{book}
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
			return
		}
	}
}



func BookHandler(c *gin.Context) {
	type Data struct{

		Book repository.Book
	}

	bookRepo := postgres.NewBooksRepository(dataBase.Connection)
	books, _ := bookRepo.GetByID(2)

	output := Data{books}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}


/*func (b BookService) getBook(c *gin.Context) {
	// Check if the bookID is valid
	if bookID, err := strconv.Atoi(c.Param("book_id"));

		err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {

		// Check if the category exists
		if book, err := b.BooksRepo.GetByID(bookID);
			err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			fmt.Print(book)
			c.JSON(http.StatusOK, book)
			return
		}
	}
}*/

func (b *BookService) insertNewBook(c *gin.Context) {
	var book repository.Book
	var err error

	if err = c.BindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	book.Base64Img =book.Base64Img[strings.IndexByte(book.Base64Img, ',')+1:]
	if  book.Image, err =base64.StdEncoding.DecodeString(book.Base64Img); err!=nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	if err := b.BooksRepo.InsertNewBook(book); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	go NotifyAllNewBook(book.Title,book.Description)
}

const preReservedTimeAllowedSecs = 1000

//PreReservedTimer is a funtion that changes the status of a book to free if user didnt propose book to exchange
//must be started from "go PreReservedTimer"
func (b * BookService)PreReservedTimer(bookId int){
	time.Sleep(preReservedTimeAllowedSecs*time.Minute)
	book,err:=b.BooksRepo.GetByID(bookId)
	if err!=nil{
		log.Println(err)
		return
	}

	if book.State == repository.BookStateReserved{
		b.BooksRepo.UpdateBookState(bookId,repository.BookStateFree)
		return
	}
	return
}
