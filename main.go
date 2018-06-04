package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/gin-gonic/gin"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"

)



func main() {
	//books.GetByCatCertainBook(2,3)
	gin.SetMode(gin.ReleaseMode)
	//Opens database connection
	connection:= dataBase.InitializeConnection()
	defer connection.Close()

	//For local testing uncomment port in init
	services.InitializeRouter()




}



