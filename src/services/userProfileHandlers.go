package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"
)


func (s *UserService)UserCreateHandler(c *gin.Context) {
	var u entity.User
	if err:=c.BindJSON(&u); err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	if err :=s.UsersRepo.InsertUser(u); err!=nil{
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	}
}

func (s *UserService)UserGetHandler(c *gin.Context) {


//	c.JSON(http.StatusOK,)
}

func (s *UserService)UserDeleteHandler(c *gin.Context) {


	c.Redirect(http.StatusTemporaryRedirect, "/")
}

/* Input example
{
	"id": 1,
	"nickname": "Denchick",
	"email": "away4ppel@den.ua",
	"password": "pass",
	"registrDate": "2018-01-01"
}
 */
func (s *UserService)UserUpdateHandler(c *gin.Context) {

}


