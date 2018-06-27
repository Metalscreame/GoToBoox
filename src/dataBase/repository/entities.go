package repository

import "time"

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
}

type User struct {
	ID                              int       `json:"-"`
	Nickname                        string    `json:"nickname"`
	Email                           string    `json:"email"`
	Password                        string    `json:"password"`
	NewPassword                     string    `json:"new_passwordd"`
	ExchangesNumber                 int       `json:"-"`
	HasBookForExchange              bool      `json:"has_book_for_exchange"`
	BookId                          int       `json:"-"`
	NotificationGetBewBooks         bool      `json:"notification_get_new_books"`
	NotificationGetWhenBookReserved bool      `json:"notification_get_when_book_reserved"`
	NotificationDaily               bool      `json:"notification_daily"`
	RegisterDate                    time.Time `json:"-"`
	ReturningBookId                 int       `json:"-"`
	TakenBookId                     int       `json:"taken_book_id"`
}

type Tags struct {
	ID    int    `json:"id,omitempty"`
	Title string `json:"tag_title"`
}

type BookTags struct {
	BookID int `json:"book_id"`
	TagID  int `json:"tag_id"`
}
