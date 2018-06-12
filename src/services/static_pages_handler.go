package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/midlwares"
)

//ShowLoginPage is a handler function that renders static login page
func ShowLoginPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "Login Page",
			"page": "login",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//ShowRegistrPage is a handler function that renders static registration page
func ShowRegistrPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "Registration Page",
			"page": "registration",
			"isLoggedIn": isLoggedIn,
		},
	)
}

//UserProfileHandler is a handler func that handle /userProfile handler and decides whether user is logged in or not
//If not, it redirects to login page, else - to the usersProfilePage
func  UserProfileHandler(c *gin.Context) {
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
func(s* UserService) ShowUsersProfilePage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"index.tmpl.html",
		gin.H{
			"title": "User's profile page",
			"page": "userprofile",
			"isLoggedIn": isLoggedIn,
		},
	)
}

func ShowBook(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(http.StatusOK, "index.tmpl.html", gin.H{
		"title": "Book - Description",
		"page" : "book",
		"isLoggedIn": isLoggedIn,

	})
}

func ShowUploadBookPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"uploadBookPage.html",
		gin.H{
			"title": "Upload Book Page",
			"page": "uploadpage",
			"isLoggedIn": isLoggedIn,
		},
	)
}

func ShowTakenBooksPage(c *gin.Context) {
	isLoggedIn := midlwares.CheckLoggedIn(c)
	c.HTML(
		http.StatusOK,
		"takenBooksPage.html",
		gin.H{
			"title": "Taken books",
			"page": "takenBooks",
			"isLoggedIn": isLoggedIn,
		},
	)
}