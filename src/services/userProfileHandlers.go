package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
)

/* Input example for create and update
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01"
}
 */
func (s *UserService)UserCreateHandler(c *gin.Context) {
	var u entity.User
	if err:=c.BindJSON(&u); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err :=s.UsersRepo.InsertUser(u); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusTemporaryRedirect, gin.H{"status": "ok"})
	c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (s *UserService)UserGetHandler(c *gin.Context) {


//	c.JSON(http.StatusOK,)
}

func (s *UserService)UserDeleteHandler(c *gin.Context) {


	c.Redirect(http.StatusTemporaryRedirect, "/")
}


func (s *UserService)UserUpdateHandler(c *gin.Context) {

}


