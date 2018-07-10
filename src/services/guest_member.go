package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
	"crypto/md5"
	"encoding/hex"
	"time"
	"gopkg.in/appleboy/gin-jwt.v2"
	gojwt "github.com/dgrijalva/jwt-go"
)

// UserService is a struct that is is used to set repository for usersRepo (or its mocks)
type UserService struct {
	UsersRepo repository.UserRepository
}

//NewUserService is a func to get new UserService with user's defined repository
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
		log.Println("Error in UserGetHandler while getting user from db: ")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, user)
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
		log.Println("Error in UserDeleteHandler while deleting user from db: ")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("is_logged_in", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

//UserUpdateHandler is a handler function that updates users info in database. Uses PUT method
/*Input example for update
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01",
	"role":""
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
		log.Println("Error in UserUpdateHandler while getting user from db: ")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	newPasswordToCheck := GetMD5Hash(userToUpdate.Password)
	if userFromDb.Password != newPasswordToCheck {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "passwords doesnt much"})
		return
	}
	userToUpdate.NewPassword = GetMD5Hash(userToUpdate.NewPassword)

	if err := s.UsersRepo.UpdateUserByEmail(userToUpdate, email); err != nil {
		log.Println("Error in UserUpdateHandler while updating user in db: ")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}
	c.SetCookie("email", userToUpdate.Email, 2*60*60, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

//LogoutHandler is a handler function that logging out from site and clears users cookie
//Uses route /api/v1/logout
func (s *UserService) LogoutHandler(c *gin.Context) {
	c.SetCookie("email", "", -1, "", "", false, true)
	c.SetCookie("nickname", "", -1, "", "", false, true)
	c.SetCookie("token", "", -1, "", "", false, true)
	c.SetCookie("is_logged_in", "", -1, "", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(http.StatusFound, "/")
}

//UserCreateHandler is a handler function that creates new user in a database and generate session token
/*Uses route/api/v1/register
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
	var tokenString string

	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	u.RegisterDate = time.Now().UTC()
	u.Password = GetMD5Hash(u.Password)
	lastID, err := s.UsersRepo.InsertUser(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	if err = s.UsersRepo.InsertRolesToUsers(lastID, 1); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	token := gojwt.New(gojwt.SigningMethodHS256)
	token = gojwt.NewWithClaims(gojwt.SigningMethodHS256, gojwt.MapClaims{
		"id":            string(lastID),
		"role":          "user",
		"exp":           time.Now().Add(time.Hour * 2).Unix(),
		"generatedData": time.Now(),
	})

	tokenString, err = token.SignedString(jwtMiddleware.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "bad request"})
		return
	}
	c.SetCookie("token", tokenString, 2*60*60, "", "", false, false)
	performLoginCookiesSetting(u, c)
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
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
}

//CheckCredentials checks email and password
func (s *UserService) CheckCredentials(email string, password string, c *gin.Context) (string, bool) {

	user, err := s.UsersRepo.GetUserByEmail(email)
	if err != nil || user.Password != GetMD5Hash(password) {
		return "", false
	}
	s.Payload(email)
	performLoginCookiesSetting(user, c)
	return email, true
}

//Authorization checks user's permission to access special links
func (s *UserService) Authorization(userID string, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if claims["role"] == "admin" {
		return true
	}
	return false
}

//Payload is a function that add some fields to the jwt
func (s *UserService) Payload(email string) (map[string]interface{}) {
	b := s.UsersRepo
	user, _ := b.GetRoleByEmail(email)

	return map[string]interface{}{
		"id":            user.ID,
		"role":          user.Role,
		"exp":           time.Now().Add(time.Hour * 2).Unix(),
		"generatedData": time.Now(),
	}
}

func isUserValid(email string, password string, repository repository.UserRepository) bool {
	user, err := repository.GetUserByEmail(email)

	if err != nil || user.Password != GetMD5Hash(password) {
		return false
	}
	return true
}

func performLoginCookiesSetting(u repository.User, c *gin.Context) {

	c.Set("is_logged_in", true)
	c.SetCookie("email", u.Email, 2*60*60, "", "", false, false)
	c.SetCookie("nickname", u.Nickname, 2*60*60, "", "", false, false)
	c.SetCookie("is_logged_in", "true", 2*60*60, "", "", false, false)
}

//GetMD5Hash generates md5 hash from input string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
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
