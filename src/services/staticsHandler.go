package services


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login Page",
		},
	)
}

func ShowRegistrPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"registration.html",
		gin.H{
			"title": "Registration Page",
		},
	)
}

func ShowUsersProfilePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"userProfile.html",
		gin.H{
			"title": "User's profile",
		},
	)
}
