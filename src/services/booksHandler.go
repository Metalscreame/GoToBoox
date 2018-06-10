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
)

type BookService struct {
	BooksRepo repository.BookRepository
	UsersRepo repository.UserRepository
}

func NewBookService(repository repository.BookRepository, usersRepo repository.UserRepository) *BookService {
	return &BookService{
		BooksRepo: repository,
		UsersRepo:usersRepo,
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

	reservedBookId, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if reservedBookId != 0{
		emailCookie, err := c.Request.Cookie("email")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
			return
		}
		err=b.BooksRepo.UpdateBookStateAndUsersBookIdByUserEmail(emailCookie.String(),repository.BookStateReserved, reservedBookId)
		if err != nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
	}

	if err := b.BooksRepo.InsertNewBook(book); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	go NotifyAllNewBook(book.Title,book.Description)
	go b.ReservedTimer(reservedBookId)
	return
}

//UpdateBookStatusToReturningFromTaken is a handler func that sets up state from taken to returned if user want this book
func(b *BookService) UpdateBookStatusToReturningFromTaken(c *gin.Context) {

	bookID,err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"bad request"})
		return
	}

	err=b.BooksRepo.UpdateBookState(bookID,repository.BookStateReturningToShelf)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

//ExchangeBook is a handler func that make a bookcross feature, sets all the values and clears everything
func (b *BookService) ExchangeBook(c * gin.Context)  {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	email := emailCookie.String()

	err=b.UsersRepo.MakeBookCross(email)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	user,err:=b.UsersRepo.GetUserByEmail(email)
	if err!=nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	if user.Returning_book_id>0{
		err=b.UsersRepo.ClearReturningBookIdByEmail(email)
		if err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
		err=b.BooksRepo.UpdateBookState(user.Returning_book_id,repository.BookStateFree)
		if err!=nil{
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
	}
	return
}

//UpdateBookStatusToReturning is a handler func that handled route of taken books that wants to be back, when user
//reserves s new book but have previoysly registered at the system one
func(b *BookService) UpdateBookStatusToReturning(c *gin.Context) {
	takenBookId,err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	reservedBookId,err := strconv.Atoi(c.Param("reserved_book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	email:=emailCookie.String()

	err=b.BooksRepo.UpdateBookStateAndUsersBookIdByUserEmail(email,repository.BookStateReserved, reservedBookId)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	err=b.BooksRepo.UpdateBookState(takenBookId,repository.BookStateReturningToShelf)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	err= b.UsersRepo.SetReturningBookIdByEmail(takenBookId,email)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	book,err:=b.BooksRepo.GetByID(reservedBookId)
	if err != nil{
		log.Println(err)
		return
	}
	go b.ReservedTimer(reservedBookId)
	go NofityAllBookReserved(book.Title,book.Description)
	return
}


const preReservedTimeAllowedSecs = 1000

//PreReservedTimer is a funtion that changes the status of a book to free if user didnt propose book to exchange
//must be started from "go PreReservedTimer"
func (b * BookService)ReservedTimer(bookId int){
	if bookId!=0{
		time.Sleep(preReservedTimeAllowedSecs*time.Minute)
		book,err:=b.BooksRepo.GetByID(bookId)
		if err!=nil{
			log.Println(err)
			return
		}

		if book.State == repository.BookStateReserved{
			err = b.BooksRepo.UpdateBookState(bookId,repository.BookStateFree)
			if err!=nil{
				log.Println(err)
				return
			}
			err=b.UsersRepo.SetUsersBookAsNullByBookId(bookId)
			if err!=nil{
				log.Println(err)
				return
			}
			return
		}
	}
	return
}
