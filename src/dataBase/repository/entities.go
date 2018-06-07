package repository

import "time"

type User struct {
	ID           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	RegisterDate time.Time `json:"-"`
	RegTimeStr   string    `json:"registrDate"`
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

type BookDescription struct {
	BookTitle     string  `json:"bookTitle"`
	Description   string  `json:"description"`
	Popularity    float32 `json:"popularity"`
	CategoryTitle string  `json:"CategoryTitle"`
}
