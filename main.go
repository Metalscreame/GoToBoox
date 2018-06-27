package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"
	"os"
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
)

func main() {
	//os.Setenv("POSTGRES_URL","postgres://niuaznefoznzqh:33f9db3d3723c0a337c18e6f0c599d358765159048ab0c4ec5a1d28969616854@ec2-54-217-214-68.eu-west-1.compute.amazonaws.com:5432/dcclgdtqr61bti")
	//os.Setenv("PORT","8080")
	file:=setupLogFile()
	defer file.Close()
	dataBase.Connect()
	postgres.NewBooksRepository(dataBase.Connection).GetByID(1)
	services.Start()
	go services.DailyEmailNotifications()
}

func setupLogFile()  *os.File{
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	println("All errors will be in the log.txt. Read it if you think that something is wrong.")
	log.Println("Recording of the log file has started...")
	return logFile
}
