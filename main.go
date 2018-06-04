package main

import (
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/metalscreame/GoToBoox/src/dataBase"
	"github.com/metalscreame/GoToBoox/src/services"

	"os"
	"log"
	"bufio"
)

//envVariable is a variable that stores run mode for server. if its "production" than its a heroku server, and we need
//to start it in production mode. If its empty = its a local machine with no such variable.
const envVariable = "GOLANG_RUN_MODE"

func main() {
	credentials:=getDatabaseCredentials()
	dataBase.Connect(credentials)
	services.Start()
}


func getDatabaseCredentials() (dbConf dataBase.DataCredentials) {
	var file *os.File
	var err error

	runMode := os.Getenv(envVariable)
	if runMode == "production" {
		file, err = os.Open("productionConfig")
		CheckForFatalError(err)
		dbConf=readConfigValuesFromFile(file)
	} else {
		file, err = os.Open("developmentConfig")
		CheckForFatalError(err)
		dbConf=readConfigValuesFromFile(file)
	}
	if err := file.Close(); err != nil {
			panic(err)
	}
	return
}

func readConfigValuesFromFile(file *os.File)(dbConf dataBase.DataCredentials){
	buf := bufio.NewReader(file)
	dbConf.DB_NAME = getSingleLineFromFile(buf)
	dbConf.DB_PASSWORD = getSingleLineFromFile(buf)
	dbConf.DB_USER = getSingleLineFromFile(buf)
	return
}

func getSingleLineFromFile(r *bufio.Reader)  string {
	a,err := r.ReadString('\n')
	CheckForFatalError(err)
	return string(a[:len(a)-2])
}

//CheckForFatalError is an error handler function that stops program when a serious error occur
func CheckForFatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

