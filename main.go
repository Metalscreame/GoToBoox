package main

import (
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"

	"os"
	"log"
	"io/ioutil"
	"encoding/json"

)

//envVariable is a variable that stores run mode for server. if its "production" than its a heroku server, and we need
//to start it in production mode. If its empty = its a local machine with no such variable.
const envVariable = "GOLANG_RUN_MODE"

func main() {
	file :=setupLogFile()
	defer file.Close()

	credentials, port := getDatabaseCredentialsAndPort()
	dataBase.Connect(credentials)
	services.Start(port)

}

func getDatabaseCredentialsAndPort() (d dataBase.DataBaseCredentials, port string) {
	runMode := os.Getenv(envVariable)
	if runMode == "production" {
		bytes, err := ioutil.ReadFile("productionConfig.json")
		CheckForFatalError(err)
		d = readConfigValuesFromFile(bytes)
		port = os.Getenv("PORT")
	} else {
		bytes, err := ioutil.ReadFile("developmentConfig.json")
		CheckForFatalError(err)
		d = readConfigValuesFromFile(bytes)
		port = "8080"
	}
	return
}

func readConfigValuesFromFile(b []byte) (d dataBase.DataBaseCredentials) {
	err := json.Unmarshal(b, &d)
	CheckForFatalError(err)
	return
}

func setupLogFile()  *os.File{
	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0666)
	CheckForFatalError(err)
	log.SetOutput(logFile)
	log.Println("Recording of the log file has started...")
	return logFile
}

//CheckForFatalError is an error handler function that stops program when a serious error occur
func CheckForFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
