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
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"registration.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Registration Page",
		},
	)
}

func ShowUsersProfilePage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,
		// Use the index.html template
		"userProfile.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "User's profile",
		},
	)
}
