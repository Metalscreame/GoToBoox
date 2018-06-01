package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/services"
	db "github.com/metalscreame/GoToBoox/src/dataBase/dbConnection"
)


func main() {

	//Opens database connection
	db.InitializeConnection()
	defer db.GlobalDataBaseConnection.Close()

	//For local testing uncomment port in init
	services.InitializeRouter()
	//models.GetCategories()
}
