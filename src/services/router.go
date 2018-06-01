package services

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	apiV1route       = "/api/v1"
	userProfileRoute = "/userProfile"
)


var router *gin.Engine

func InitializeRouter() {
	//Used for heroku
	port := os.Getenv("PORT")

	//Uncomment for local machine   !!!!
	//port="8080"


	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router = gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/api/v.1/", api.Handler)
	//The place for handlers routes
	// exm router.GET("/", showIndex.ShowIndexPage)
	initUserProfileRouters()
	router.Run(":" + port)
}

func initUserProfileRouters(){
	router.GET(apiV1route+userProfileRoute,userProfile.UserGet)
	router.POST(apiV1route+userProfileRoute,userProfile.UserCreate)
	router.PUT(apiV1route+userProfileRoute,userProfile.UserUpdate)
	router.DELETE(apiV1route+userProfileRoute,userProfile.UserDelete)
}