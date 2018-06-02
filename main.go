package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/services"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)



func main() {

	//Opens database connection
	connection:= dataBase.InitializeConnection()
	defer connection.Close()

	//For local testing uncomment port in init
	services.InitializeRouter()
	//models.GetCategories()
}


