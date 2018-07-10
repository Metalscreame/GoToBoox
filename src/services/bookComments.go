package services

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
	"time"
)

//CommentsService is a struct that holds Comments repository
type CommentsService struct {
	CommentsRepo repository.CommentsRepository
}

//NewCommentsService is a function to return comments service based on concrete repository
func NewCommentsService(commentsRepo repository.CommentsRepository) *CommentsService {
	return &CommentsService{
		CommentsRepo: commentsRepo,
	}
}
//Data is a structure that holds comments
type Data struct {
	Comments []repository.Comment
}

//BookCommentsHandler is a handler func that returns all the comments for a single book.
func (cs *CommentsService) BookCommentsHandler(c *gin.Context) {
	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	comments, _ := cs.CommentsRepo.GetAllCommentsByBookID(bookID)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Data{comments}})

}

//AddBookCommentHandler is a handler func that adds a single user's comment for a single book.
func (cs *CommentsService) AddBookCommentHandler(c *gin.Context) {

	var comment repository.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
		return
	}

	bookID, err := strconv.Atoi(c.Param("book_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	emailCookie, err := c.Request.Cookie("email")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}
	nicknameCookie, err := c.Request.Cookie("nickname")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "bad request"})
		return
	}

	nickname := nicknameCookie.Value
	email := convertEmailString(emailCookie.Value)

	err = cs.CommentsRepo.InsertNewComment(email, nickname, comment.CommentaryText, bookID)
	if err != nil {
		log.Println("Error in AddBookCommentHandler while getting user from db: ")
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}


//AllCommentsByNicknameHandler is a handler func, that returns all the comments that user wrote.
func (cs *CommentsService) AllCommentsByNicknameHandler(c *gin.Context)  {
	nickname := c.Param("nickname")

	comments,err:=cs.CommentsRepo.GetAllCommentsByNickname(nickname)
	if err != nil {
		log.Printf("Error in AllCommentsByNickname while getting user from db at %v: ",time.Now())
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": Data{comments}})
}