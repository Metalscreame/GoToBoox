package repository

import "time"

type User struct {
	ID           int       `json:"-"`
	Nickname     string    `json:"nickname"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	NewPassword  string    `json:"new_passwordd"`
	RegisterDate time.Time `json:"-"`
}

type Categories struct {
	ID    int
	Title string
}

type Book struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	Popularity   float32 `json:"popularity"`
	CategoriesID int     `json:"categoriesID"`
}

type Authors struct {
	ID         int
	FirstName  string
	MiddleName string
	LastName   string
}

type UsersBooks struct {
	user User
	Books []Book
}

