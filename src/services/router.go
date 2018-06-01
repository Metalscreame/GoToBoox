package services

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/api/v.1"
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

	//The place for api-handlers routes
	router.GET("/api/v.1/", api.Handler)


	// The place for ui-handlers routes
	// router.GET("/", ui.Index)




	router.Run(":" + port)
}