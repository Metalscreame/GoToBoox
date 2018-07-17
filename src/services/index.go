package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

// IndexHandler get all nedeed data for the main page from repos.
func IndexHandler(c *gin.Context) {
	type Data struct{
		Books []repository.Book
	}

	bookRepo := postgres.NewBooksRepository(dataBase.Connection)
	books, _ := bookRepo.GetAll()

	output := Data{books}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

