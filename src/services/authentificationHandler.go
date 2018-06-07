package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"math/rand"
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"crypto/md5"
	"encoding/hex"
)

//LogoutHandler is a handler function that logging out from site and clears users cookie
//Uses route /api/v1/logout
func (s *UserService) LogoutHandler(c *gin.Context) {
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(http.StatusFound, "/")
	return
}

/* UserCreateHandler is a handler function that creates new user in a database\
Uses route/api/v1/register
Input example for create
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01"
}
 */
func (s *UserService) UserCreateHandler(c *gin.Context) {
	var u repository.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	u.RegisterDate = time.Now()
	u.Password = getMD5Hash(u.Password)

	if err := s.UsersRepo.InsertUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	performLoginCookiesSetting(u, c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

//PerformLoginHandler is a handler to handle loggining and setting cookies after success login
//Uses route /api/v1/login
func (s *UserService) PerformLoginHandler(c *gin.Context) {
	var u repository.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	if isUserValid(u.Email, u.Password, s.UsersRepo) {
		performLoginCookiesSetting(u, c)
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"status": "wrong credentials"})
	return
}

func isUserValid(email string, password string, repository repository.UserRepository) bool {
	user, err := repository.GetUserByEmail(email)
	if err != nil || user.Password != password {
		return false
	}
	return true
}

// I'm using a random 16 character string as the session token
// This is not a secure way of generating session tokens
func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}

func performLoginCookiesSetting(u repository.User, c *gin.Context) {
	token := generateSessionToken()
	c.SetCookie("token", token, 16000, "", "", false, true)
	c.Set("is_logged_in", true)
	c.SetCookie("email", u.Email, 16000, "", "", false, true)
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
