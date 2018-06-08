package repository

type BookRepository interface {
	GetAll() ([]interface{}, error)
	GetByCategory(categoryID int) ([]Book, error)
	GetByID(bookID int) (b Book, err error)
	GetMostPopularBooks(quantity int) ([]Book, error)
}

