package repository

import "time"

//Categories is a entity struct for Categories table
type Categories struct {
	ID    int
	Title string
}

const (
	BookStateFree             = "FREE"
	BookStateReserved         = "RESERVED"
	BookStateTaken            = "TAKEN"
	BookStateReturningToShelf = "RETURNING"
)

//Book is a entity struct for Book table
type Book struct {
	ID             int      `json:"id,omitempty"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Popularity     float32  `json:"popularity,omitempty"`
	EvaluateNumber int      `json:"-"`
	State          string   `json:"state,omitempty"`
	Image          []byte   `json:"image,omitempty"`
	Base64Img      string   `json:"base_64_img"`
	TagID          []string `json:"tag_id"`
	TagsTitles     string   `json:"tag_title"`
}

//User is a entity struct for User table
type User struct {
	ID                              int       `json:"-"`
	Nickname                        string    `json:"nickname" `
	Email                           string    `json:"email" binding:"required"`
	Password                        string    `json:"password" binding:"required"`
	NewPassword                     string    `json:"new_passwordd"`
	ExchangesNumber                 int       `json:"-"`
	HasBookForExchange              bool      `json:"has_book_for_exchange"`
	BookID                          int       `json:"-"`
	NotificationGetBewBooks         bool      `json:"notification_get_new_books"`
	NotificationGetWhenBookReserved bool      `json:"notification_get_when_book_reserved"`
	NotificationDaily               bool      `json:"notification_daily"`
	RegisterDate                    time.Time `json:"-"`
	ReturningBookID                 int       `json:"-"`
	TakenBookID                     int       `json:"taken_book_id"`
	Role                            string    `json:"role"`
}

//Tags is a entity struct for Tags table
type Tags struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"tag_title"`
}

//BookTags is a entity struct for comments BookTags table
type BookTags struct {
	BookID int `json:"book_id"`
	TagID  int `json:"tag_id"`
}

//Comment is a entity struct for comments table
type Comment struct {
	ID             int       `json:"-"`
	BookID         int       `json:"book_id,omitempty"`
	UserNickname   string    `json:"nickname,omitempty"`
	UserEmail      string    `json:"-"`
	CommentaryText string    `json:"commentText"`
	CommentDate    time.Time `json:"-"`
	FormatedDate   string    `json:"date,omitempty"`
}
