package services

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/books"
	"net/http"
)


func (b *BookService) showAllBooks(c *gin.Context) {
	books, err := books.GetAll()
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {

	// Call the render function with the name of the template to render
	c.JSON(http.StatusOK, books)
	}
}

//Get all books by certain category
func (b *BookService) getBooks(c *gin.Context) {
	// Check if the categoryID is valid
	if catID, err := strconv.Atoi(c.Param("cat_id"));
	err == nil {
		// Check if the category exists
		if book, err := books.GetByCategory(catID);
		err == nil {
			c.JSON(http.StatusOK, book)
			return

		} else {
			// If the book is not found, abort with an error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	} else {
		// If an invalid category ID is specified in the URL, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func (b *BookService) getBook (c *gin.Context) {
	// Check if the bookID is valid
	if bookID, err := strconv.Atoi(c.Param("book_id"));
	err == nil {
		// Check if the category exists
		if book, err := books.GetByID(bookID); err == nil {

			c.JSON(http.StatusOK, book)
			return

		} else {
			// If the book is not found, abort with an error
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	} else {
		// If an invalid book ID is specified in the URL, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

/*func (b * BookService) getByCatCertainBook (c *gin.Context){
	if catID, err := strconv.Atoi(c.Param("cat_id"));
		err == nil {
		if bookID, err := strconv.Atoi(c.Param("book_id"));
		err == nil{

			if book, err := books.GetByCatCertainBook(catID,bookID); err == nil {
				c.JSON(http.StatusOK, book)
				return
		}
		}
	} else {
		// If an invalid book ID is specified in the URL, abort with an error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
		}
	}*/



