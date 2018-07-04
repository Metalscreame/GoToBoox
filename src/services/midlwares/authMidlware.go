package midlwares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// This middleware ensures that a request will be aborted with an error
// if the user is not logged in
func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.Redirect(http.StatusFound, "/")
		}
	}
}


// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if loggedIn {
			c.Redirect(http.StatusFound, "/")
		}
		}
	}


// This middleware sets whether the user is logged in or not
func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Request.Cookie("token"); err == nil || token.String() != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}

func CheckLoggedIn(c *gin.Context) bool {
	_, err := c.Request.Cookie("is_logged_in")
	if err != nil {
		return false
	}
	return true
}
