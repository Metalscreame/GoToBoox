package main

import (
	_ "gopkg.in/github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"
	"os"
	"log"
	"github.com/metalscreame/GoToBoox/src/dataBase/postgres"
)

func main() {
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
