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
	credentials, port := getDatabaseCredentialsAndPort()
	dataBase.Connect(credentials)
	services.Start(port)
}

func getDatabaseCredentialsAndPort() (d dataBase.DataBaseCredentials, port string) {
	runMode := os.Getenv(envVariable)
	if runMode == "production" {
		bytes, err := ioutil.ReadFile("productionConfig")
		CheckForFatalError(err)
		d = readConfigValuesFromFile(bytes)
		port = os.Getenv("PORT")
	} else {
		bytes, err := ioutil.ReadFile("developmentConfig")
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

//CheckForFatalError is an error handler function that stops program when a serious error occur
func CheckForFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
