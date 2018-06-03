package entity

import "time"

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RegistrDate time.Time `json:"-"`
	RegTimeStr string `json:"registrDate"`
}

type Categories struct {
	Id    int
	Title string `json:"categories"`
}

type Book struct {
	ID    int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Popularity float32 `json:"popularity"`
	CategoriesID int `json:"categoriesID"`
}

type Authors struct {
	Id        int
	FirstName string
	MidleName string
	LastName  string
}

type BooksAuthors struct {
	BookID   int
	AuthorID int
}
