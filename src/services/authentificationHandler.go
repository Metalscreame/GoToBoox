package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"math/rand"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/users"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
)

/*
	IMPORTANT COMMENT section
	Writing this i found a bug. When u do redirect from post method -  it gives 404
	Researching told me that solves this http.StatusFound

	Yours Roman Kosyiy
 */



//LogoutHandler is a handler function that logging out from site and clears users cookie
// 				/api/v1/logout
func (s *UserService) LogoutHandler(c *gin.Context) {
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.Redirect(http.StatusOK, "/")
}

/* UserCreateHandler is a handler function that creates new user in a database\
/api/v1/register
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
	var u entity.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.UsersRepo.InsertUser(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	performLoginCookiesSetting(u,c)
	c.JSON(http.StatusFound, gin.H{"status": "ok"})
	c.Redirect(http.StatusFound, "/")
}

//PerformLoginHandler is a handler to handle loggining and setting cookies after success login
// /api/v1/login
func (s *UserService) PerformLoginHandler(c *gin.Context) {
	var u entity.User

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if isUserValid(u.Email, u.Password, s.UsersRepo) {
		performLoginCookiesSetting(u,c)
		c.Redirect(http.StatusFound, "/")
	} else {
		c.Redirect(http.StatusFound, "/login")
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

func performLoginCookiesSetting(u entity.User,c *gin.Context) {
	token := generateSessionToken()
	c.SetCookie("token", token, 16000, "", "", false, true)
	c.Set("is_logged_in", true)
	c.SetCookie("email", u.Email, 16000, "", "", false, true)
}
