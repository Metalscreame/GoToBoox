package repository

type CommentsRepository interface {
	GetAllCommentsByNickname(nickname string) (comments []Comment, err error)
	GetAllCommentsByBookID(bookID int) (comments []Comment,err error)
	InsertNewComment(email,nickname,comment string,bookID int) (err error)
}
