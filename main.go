package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	_ "github.com/lib/pq"
	"github.com/metalscreame/GoToBoox/src/services"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)



func main() {
	dataBase.InitializeConnection()
	services.InitializeRouter()
}