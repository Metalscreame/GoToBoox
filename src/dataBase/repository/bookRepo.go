package repository

type BookRepository interface {
	GetAll() (books []BookDescription, err error)
	GetByCategory(categoryID int) (books []Book, err error)
	GetByID(bookID int) (books BookDescription, err error)
	GetMostPopularBooks(quantity int) ([]Book, error)
	}