package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/services"
)

func main() {
	//For local testing uncomment port in init
	services.InitializeRouter()
}
