package repository

type BookRepository interface {
	GetAll() (books []BookDescription, err error)
	GetByCategory(categoryID int) (books []Book, err error)
	GetByID(bookID int) (book Book, err error)
	GetMostPopularBooks(id int) ([]Book, error)
}
