package services


import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLoginPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,

		"login.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "Login Page",
		},
	)
}

func ShowRegistrPage(c *gin.Context) {
	c.HTML(
		// Set the HTTP status to 200 (OK)
		http.StatusOK,

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

		"usersProfile.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "User's profile",
		},
	)
}
