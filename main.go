package main

import (
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"
	"os"
	"log"
)

func main() {
	file:=setupLogFile()
	defer file.Close()
	dataBase.Connect()
	services.Start()
}

func setupLogFile()  *os.File{
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logFile)
	log.Println("Recording of the log file has started...")
	return logFile
}
