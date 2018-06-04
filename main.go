package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/services"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)



func main() {
	dataBase.InitializeConnection()
	services.InitializeRouter()
}