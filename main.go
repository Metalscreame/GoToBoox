package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/gin-gonic/gin"

)



func main() {
	gin.SetMode(gin.ReleaseMode)
	//Opens database connection
	//connection:= dataBase.InitializeConnection()
	//defer connection.Close()

	//For local testing uncomment port in init
	//services.InitializeRouter()
	//models.GetCategories()



}



