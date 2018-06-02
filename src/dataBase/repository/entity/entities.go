package entity

type User struct {
	Id          int    `json:"id"`
	Nickname    string `json:"nickname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	RegistrDate string `json:"registr_date"`
}

type Categories struct {
	Id    int
	Title string
}

type Book struct {
	Id           int
	Title        string
	Description  string
	Popularity   float32
	CategoriesID int
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