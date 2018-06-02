package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/users"
)

func (s *UserService) LogoutHandler(c *gin.Context) {
	c.SetCookie("nickname", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (s *UserService) PerformLoginHandler(c *gin.Context) {
	var u entity.User
	c.BindJSON(&u)

	if isUserValid(u.Email, u.Password, s.UsersRepo) {
		token := generateSessionToken()
		c.SetCookie("token", token, 16000, "", "", false, true)
		c.Set("is_logged_in", true)
		c.SetCookie("nickname", "", 16000, "", "", false, true)
		c.Redirect(http.StatusTemporaryRedirect, "/")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func isUserValid(email string, password string, repository users.UserRepository) bool {
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
