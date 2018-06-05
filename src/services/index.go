package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func IndexHandler(c *gin.Context) {
	type Data struct{
		PopularBooks []repository.Book
		Categories []repository.Categories
	}

	bookRepo := postgres.NewBooksRepository(dataBase.Connection)
	books, _ := bookRepo.GetMostPopularBooks(5)
	catRepo := postgres.CategoryRepoPq{}
	cats, _ := catRepo.GetAllCategories()

	output := Data{cats}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}
