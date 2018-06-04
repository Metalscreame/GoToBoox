package main

import (
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"

)



func main() {
	dataBase.InitializeConnection()
	services.InitializeRouter()
}

