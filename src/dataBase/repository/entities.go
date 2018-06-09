package repository

import "time"

type Categories struct {
	ID    int
	Title string
}

type Book struct {
	ID             int     `json:"id"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Popularity     float32 `json:"popularity"`
	EvaluateNumber int     `json:"-"`
	State          string  `json:"state"`
	Image          []byte  `json:"image"`
}

type User struct {
	ID                              int       `json:"-"`
	Nickname                        string    `json:"nickname"`
	Email                           string    `json:"email"`
	Password                        string    `json:"password"`
	NewPassword                     string    `json:"new_passwordd"`
	ExchangesNumber                 int       `json:"-"`
	HasBookForExchange              bool      `json:"has_book_for_exchange"`
	Book                            Book      `json:"-"`
	NotificationGetBewBooks         bool      `json:"notification_get_new_books"`
	NotificationGetWhenBookReserved bool      `json:"notification_get_when_book_reserved"`
	RegisterDate                    time.Time `json:"-"`
}

type Authors struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
}

type BookDescription struct {
	BookTitle     string  `json:"bookTitle"`
	Description   string  `json:"description"`
	Popularity    float32 `json:"popularity"`
	CategoryTitle string  `json:"CategoryTitle"`
}
