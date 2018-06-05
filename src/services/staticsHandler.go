package services


import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/services/authentification/midlware"
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

func UserProfileHandler(c *gin.Context) {
	loggedIn:=midlware.CheckLoggedIn(c)
	if loggedIn{
		c.Redirect(http.StatusFound,"/userProfilePage")
		return
	}else{
		c.Redirect(http.StatusFound,"/login")
		return
	}

}

func ShowUsersProfilePage(c *gin.Context)  {
	c.HTML(
		http.StatusOK,
		"userProfile.html",
		gin.H{
			"title": "Registration Page",
		},
	)
}
