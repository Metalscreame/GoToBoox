package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

func IndexHandler(c *gin.Context) {
	type Data struct{
		Books []repository.Book
		Users []repository.User
	}

	bookRepo := postgres.NewBooksRepository(dataBase.Connection)
	books, _ := bookRepo.GetAll()

	userRepo := postgres.NewPostgresUsersRepo(dataBase.Connection)
	users, _ := userRepo.GetAllUsers()

	output := Data{books, users}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": output})
}

