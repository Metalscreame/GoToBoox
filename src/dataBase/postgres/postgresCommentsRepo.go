package postgres

import (
	"database/sql"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"time"
)

type poastgresCommentRepository struct {
	Db *sql.DB
}

//NewCommentsRepository is a function to get New poastgresCommentRepository which uses given connection
func NewCommentsRepository(Db *sql.DB) repository.CommentsRepository {
	return &poastgresCommentRepository{Db}
}

func (p *poastgresCommentRepository) GetAllCommentsByNickname(nickname string) (comments []repository.Comment, err error) {
	rows, err := p.Db.Query("SELECT book_id, user_nickname, commentary, commentary_date  FROM gotoboox.comments where user_nickname = $1", nickname)
	if err != nil {
		return
	}
	defer rows.Close()

	var comment repository.Comment
	for rows.Next() {

		if err = rows.Scan(&comment.BookID, &comment.UserNickname, &comment.CommentaryText, &comment.CommentDate);
			err != nil {
			return
		}
		comment.FormatedDate = comment.CommentDate.Format("2006-01-02 15:04:05")
		comments = append(comments, comment)
	}
	return
}

func (p *poastgresCommentRepository) GetAllCommentsByBookID(bookID int) (comments []repository.Comment, err error) {
	rows, err := p.Db.Query("SELECT user_nickname, commentary, commentary_date  FROM gotoboox.comments where book_id = $1", bookID)
	if err != nil {
		return
	}
	defer rows.Close()

	var comment repository.Comment
	for rows.Next() {

		if err = rows.Scan(&comment.UserNickname, &comment.CommentaryText, &comment.CommentDate);
			err != nil {
			return
		}
		comment.BookID = bookID
		comment.FormatedDate = comment.CommentDate.Format("2006-01-02 15:04:05")
		comments = append(comments, comment)
	}

	return
}

func (p *poastgresCommentRepository) InsertNewComment(email, nickname, comment string, bookID int) (err error) {
	t := time.Now()
	_, err = p.Db.Query("INSERT INTO gotoboox.comments (user_nickname,user_email,commentary,commentary_date,book_id) values($1,$2,$3,$4,$5)",
		nickname, email, comment, t, bookID)
	return
}
