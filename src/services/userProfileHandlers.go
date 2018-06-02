package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
)

//UserGetHandler gets users data from database using unique email that is stored in cookie
//if there is no email in coolie that means that session is over
func (s *UserService) UserGetHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	user, err := s.UsersRepo.GetUserByEmail(emailCookie.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

//UserDeleteHandler deletes user from database.
func (s *UserService) UserDeleteHandler(c *gin.Context) {
	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	if err := s.UsersRepo.DeleteUserByEmail(emailCookie.String()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.LogoutHandler(c)
	return
}

/* UserUpdateHandler is a handler function that updates users info in database
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
	var u entity.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.UsersRepo.UpdateUserByEmail(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusTemporaryRedirect, gin.H{"status": "updated"})
	c.Redirect(http.StatusFound, "/")
	return
}
