package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/midlwares"
)

//ShowLoginPage is a handler function that renders static login page
func ShowLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login Page",
		},
	)
}

//ShowRegistrPage is a handler function that renders static registration page
func ShowRegistrPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"registration.html",
		gin.H{
			"title": "Registration Page",
		},
	)
}

//UserProfileHandler is a handler func that handle /userProfile handler and decides whether user is logged in or not
//If not, it redirects to login page, else - to the usersProfilePage
func UserProfileHandler(c *gin.Context) {
	loggedIn := midlwares.CheckLoggedIn(c)
	if loggedIn {
		c.Redirect(http.StatusFound, "/userProfilePage")
		return
	} else {
		c.Redirect(http.StatusFound, "/login")
		return
	}
}

//ShowUsersProfilePage is a handler function that renders static userProfile page
func ShowUsersProfilePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"userProfile.html",
		gin.H{
			"title": "Registration Page",
		},
	)
}

func ShowBook(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"book.html",
		gin.H{
			"title": "Book Description",
		},
	)
}
