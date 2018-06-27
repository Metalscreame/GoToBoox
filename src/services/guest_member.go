package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
	"strconv"
	"crypto/md5"
	"encoding/hex"
	"time"
	"math/rand"
)

type UserService struct {
	UsersRepo repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		UsersRepo: repository,
	}
}

//UserGetHandler gets users Data from database using unique email that is stored in cookie
//if there is no email in coolie that means that session is over
func (s *UserService) UserGetHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	email := convertEmailString(emailCookie.Value)
	user, err := s.UsersRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

//UserDeleteHandler deletes user from database. Uses DELETE method.
func (s *UserService) UserDeleteHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
		return
	}
	email := convertEmailString(emailCookie.Value)
	if err := s.UsersRepo.DeleteUserByEmail(email); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("is_logged_in","",-1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
	return
}

/* UserUpdateHandler is a handler function that updates users info in database. Uses PUT method
Input example for update
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01"
}
 */
func (s *UserService) UserUpdateHandler(c *gin.Context) {
	var userToUpdate repository.User
	if err := c.BindJSON(&userToUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	email := convertEmailString(emailCookie.Value)

	userFromDb, err := s.UsersRepo.GetUserByEmail(email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	userToUpdate.Password = GetMD5Hash(userToUpdate.Password)
	if userFromDb.Password != userToUpdate.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "passwords doesnt much"})
		return
	}
	userToUpdate.Password = GetMD5Hash(userToUpdate.NewPassword)

	if err := s.UsersRepo.UpdateUserByEmail(userToUpdate, email); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	c.SetCookie("email", userToUpdate.Email, 15000, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "updated"})
	return
}

//This function was created because cookies gives '%40' instead of '@' when read the email. It converts
func convertEmailString(emailCookie string) (string) {
	indexOfPercentSymb := strings.IndexRune(emailCookie, '%')
	runes := []rune(emailCookie)
	runes[indexOfPercentSymb] = '@'
	runes = append(runes[:indexOfPercentSymb+1], runes[indexOfPercentSymb+2:]...) //deletes 4
	runes = append(runes[:indexOfPercentSymb+1], runes[indexOfPercentSymb+2:]...) //deletes 0
	return string(runes)
}

//LogoutHandler is a handler function that logging out from site and clears users cookie
//Uses route /api/v1/logout
func (s *UserService) LogoutHandler(c *gin.Context) {
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("nickname", "",-1,"","",false,true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("is_logged_in", "", -1, "", "", false, true)
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
	u.Password = GetMD5Hash(u.Password)

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

	if err != nil || user.Password !=GetMD5Hash(password) {
		return false
	}
	return true
}

func performLoginCookiesSetting(u repository.User, c *gin.Context) {
	token := generateSessionToken()
	c.SetCookie("token", token, 16000, "", "", false, false)
	c.Set("is_logged_in", true)
	c.SetCookie("email", u.Email, 16000, "", "", false, false)
	c.SetCookie("nickname", u.Nickname, 16000, "", "", false, false)
	c.SetCookie("is_logged_in", "true", 16000, "", "", false, false)
}

//GetMD5Hash generates md5 hash from input string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// I'm using a random 16 character string as the session token
// This is not a secure way of generating session tokens
func generateSessionToken() string {
	return strconv.FormatInt(rand.Int63(), 16)
}
