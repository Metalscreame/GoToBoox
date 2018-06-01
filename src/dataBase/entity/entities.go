package entity

type User struct {
	Id int
	Nickname string
	Email string
	Password string
	RegistrDate string
}

type Categories struct {
	Id int
	Title string
}

type Book struct {
	Id int
	Title string
	Description string
	Popularity float32
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

