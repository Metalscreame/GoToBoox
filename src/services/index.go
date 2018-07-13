package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"github.com/metalscreame/GoToBoox/src/services/midlwares"
)

// ApiIndexHandler get all nedeed data for the main page from repos.
func ApiIndexHandler(c *gin.Context) {
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

func IndexHandler(c *gin.Context)  {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title":      "GoToBooX",
		"page":       "main",
		"isLoggedIn": isLoggedIn,
	})
}


//ServerIsOn is a function to check server status. returns 200  is server is alive.
func ServerIsOn(c * gin.Context){
	c.JSON(http.StatusOK, gin.H{"status":"alive"})
}



