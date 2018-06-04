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
		"usersProfile.html",
		// Pass the data that the page uses (in this case, 'title')
		gin.H{
			"title": "User's profile",
		},
	)
}
