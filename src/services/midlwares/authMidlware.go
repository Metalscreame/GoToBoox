package midlwares

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"gopkg.in/appleboy/gin-jwt.v2"
	gojwt "github.com/dgrijalva/jwt-go"
	"log"
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase"
)

var jwtMiddleware *jwt.GinJWTMiddleware
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

func TokenChecking() gin.HandlerFunc {
	return func(c *gin.Context) {

		if !CheckToken(c) {
			log.Printf("Wrong token was parsed by %s at %v", c.ClientIP(), time.Now().Format("2000.01.02. 01:02:03"))
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

func CheckToken(c *gin.Context) (parsed bool) {
	parsed = false
	if cookie, err := c.Request.Cookie("token"); err != nil || cookie.Value == "" {
		return false
	} else {
		tokenValue := cookie.Value
		token, _ := gojwt.Parse(tokenValue, func(token *gojwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*gojwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(dataBase.TokenKeyLookUp()), nil
		})

		if _, ok := token.Claims.(gojwt.MapClaims); ok && token.Valid {
			return true
		} else {
			return false
		}
	}
	return true

}
