package services

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
	"encoding/base64"
	"strings"
	"unicode/utf8"
	"time"
)

//BookService is a struct with book and user repository
type BookService struct {
	BooksRepo repository.BookRepository
	UsersRepo repository.UserRepository
}

//NewBookService is a func that initialize BookService struct
func NewBookService(repository repository.BookRepository, usersRepo repository.UserRepository) *BookService {
	return &BookService{
		BooksRepo: repository,
		UsersRepo: usersRepo,
	}
}

//FiveMostPop returns top 5 books
func (b *BookService) FiveMostPop(c *gin.Context) {
	FiveMostPop, _ := b.BooksRepo.GetMostPopularBooks(5)
	if len(FiveMostPop) > 0 {
		c.JSON(http.StatusOK, FiveMostPop)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "No books have been found"})
}

//showAllBooks is a handler for GetAll function
func (b *BookService) showAllBooks(c *gin.Context) {
	type Data struct {
		Book []repository.Book
	}
	books, err := b.BooksRepo.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	output := Data{books}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

func (b *BookService) showAllTakenBooks(c *gin.Context) {
	type Data struct {
		Book []repository.Book
	}
	books, err := b.BooksRepo.GetAllTakenBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	output := Data{books}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

//getBooks is a handler for GetByCategory function
func (b *BookService) getBooks(c *gin.Context) {
	// Check if the categoryID is valid
	catID, err := strconv.Atoi(c.Param("cat_id"))
	if err != nil {
		// If the book is not found, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if the category exists
	book, err := b.BooksRepo.GetByCategory(catID)
	if err != nil {
		// If an invalid category ID is specified in the URL, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

//getBook is a handler for GetByID function
func (b *BookService) getBook(c *gin.Context) {
	type Data struct {
		Book repository.Book
	}
	// Check if the bookID is valid
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := b.BooksRepo.GetByID(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output := Data{book}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

func (b *BookService) getTags(c *gin.Context) {
	type Data struct {
		Book []repository.Book
	}
	// Check if the bookID is valid
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := b.BooksRepo.GetTagsForBook(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	output := Data{book}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

func (b *BookService) getBookBySearch(c *gin.Context) {
	type Data struct {
		Books []repository.Book
	}

	// if user pass the title of the book, then we don't need use any filters - user know what book he want
	title := c.PostForm("title")
	if utf8.RuneCountInString(title) > 0 {
		books, err := b.BooksRepo.GetByLikeName(title)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		response := Data{books}
		c.JSON(http.StatusOK, gin.H{"response": response})
		return
	}

	type Myform struct {
		Tags   []string `form:"tags[]"`
		Rating []int    `form:"rating[]"`
	}
	var myform Myform
	c.Bind(&myform)
	books, err := b.BooksRepo.GetByTagsAndRating(myform.Tags, myform.Rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := Data{books}
	c.JSON(http.StatusOK, gin.H{"response": response})
}

//ShowReservedBooksByUser is a handler func that returns json with user's reserved book
func (b *BookService) ShowReservedBooksByUser(c *gin.Context) {
	type Data struct {
		Books []repository.Book
	}
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}

	email := convertEmailString(emailCookie.Value)
	u, err := b.UsersRepo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	book, _ := b.BooksRepo.GetByID(u.ReturningBookID)
	var output Data
	output.Books = append(output.Books, book)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

//ShowTakenBookByUser is a handler func that returns user's taken books
func (b *BookService) ShowTakenBookByUser(c *gin.Context) {

	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}

	email := convertEmailString(emailCookie.Value)
	u, err := b.UsersRepo.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}

	book, err := b.BooksRepo.GetByID(u.TakenBookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "no books"})
		return
	}
	book.ID = u.TakenBookID
	c.JSON(http.StatusOK, book)
}

//InsertNewBook is a handler func to add a new book to a database
func (b *BookService) InsertNewBook(c *gin.Context) {
	var bookToAdd repository.Book
	var err error

	if err = c.BindJSON(&bookToAdd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	bookToAdd.Base64Img = bookToAdd.Base64Img[strings.IndexByte(bookToAdd.Base64Img, ',')+1:]
	if bookToAdd.Image, err = base64.StdEncoding.DecodeString(bookToAdd.Base64Img); err != nil {
		log.Printf("Error in InsertBookHandler while encoding picture %v: \n", time.Now())
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	bookIDToReserve, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}

	email := convertEmailString(emailCookie.Value)
	err = b.BooksRepo.UpdateBookStateAndUsersBookIDByUserEmail(email, repository.BookStateReserved, bookIDToReserve)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	last, err := b.BooksRepo.InsertNewBook(bookToAdd)
	if err != nil {
		log.Printf("Error in InsertNewBook while adding new book to db at %v: \n", time.Now())
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	for k := range bookToAdd.TagID {
		number, _ := strconv.Atoi(bookToAdd.TagID[k])
		err = b.BooksRepo.InsertTags(number, last)
		if 	err != nil {
			log.Printf("Error in InsertNewBook while adding tag to a new book %v: \n", time.Now())
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
	}
	go NotifyAllNewBook(bookToAdd.Title, bookToAdd.Description)
}

//UpdateBookStatusToReturningFromTaken is a handler func that sets up state from taken to returned if user want this book
func (b *BookService) UpdateBookStatusToReturningFromTaken(c *gin.Context) {

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	err = b.BooksRepo.UpdateBookState(bookID, repository.BookStateReturningToShelf)
	if err != nil {
		log.Printf("Error in AllCommentsByNickname while getting user from db at %v: \n", time.Now())
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

//ExchangeBook is a handler func that make a bookcross feature, sets all the values and clears everything
func (b *BookService) ExchangeBook(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	email := convertEmailString(emailCookie.Value)

	err = b.UsersRepo.MakeBookCross(email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	user, err := b.UsersRepo.GetUserByEmail(email)
	if err != nil {
		log.Printf("Error in ExchangeBook while getting user from db %v: \n", time.Now())
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	if user.ReturningBookID != 0 {
		err = b.UsersRepo.ClearReturningBookIDByEmail(email)
		if err != nil {
			log.Printf("Error in ExchangeBook while clearing bookIdByEmail from db %v: \n", time.Now())
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
		err = b.BooksRepo.UpdateBookState(user.ReturningBookID, repository.BookStateFree)
		if err != nil {
			log.Printf("Error in ExchangeBook while updating book state in db %v: \n", time.Now())
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
			return
		}
	}
}

//UpdateBookStatusToReturning is a handler func that handled route of taken books that wants to be back, when user
//reserves s new book but have previoysly registered at the system one
func (b *BookService) UpdateBookStatusToReturning(c *gin.Context) {
	takenBookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	reservedBookID, err := strconv.Atoi(c.Param("reserved_book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	email := convertEmailString(emailCookie.Value)

	err = b.BooksRepo.UpdateBookStateAndUsersBookIDByUserEmail(email, repository.BookStateReserved, reservedBookID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	err = b.BooksRepo.UpdateBookState(takenBookID, repository.BookStateReturningToShelf)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	err = b.UsersRepo.SetReturningBookIDByEmail(takenBookID, email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}

	c.Redirect(http.StatusFound, "/")
	book, err := b.BooksRepo.GetByID(reservedBookID)
	if err != nil {
		log.Println(err)
		return
	}
	go NofityAllBookReserved(book.Title, book.Description)
}
