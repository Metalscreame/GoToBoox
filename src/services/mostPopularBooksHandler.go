package services

import (
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/books"
	"net/http"
)

func (bs *BookService) FiveMostPop (c *gin.Context) {
	FiveMostPop, _ := (books.BooksRepositoryPG{}).GetMostPopularBooks(5)
	if len(FiveMostPop) > 0 {
		c.JSON(http.StatusOK, FiveMostPop)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "No books have been found"} )
	}
}

