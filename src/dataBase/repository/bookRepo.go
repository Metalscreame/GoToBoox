package repository

//BookRepository is a repository interface for operations with books
type BookRepository interface {
	InsertTags (tagID int, bookID int) (err error)
	GetAll() (books []Book, err error)
	GetByCategory(categoryID int) (books []Book, err error)
	GetByLikeName(title string) (books []Book, err error)
	GetByTagsAndRating(tags []string, rating []int) (books []Book, err error)
	GetByID(bookID int) (books Book, err error)
	GetMostPopularBooks(quantity int) ([]Book, error)
	GetAllTakenBooks() (books []Book, err error)

	InsertNewBook(b Book) (lastID int, err error)
	UpdateBookState(bookID int, state string) (err error)
	UpdateBookStateAndUsersBookIDByUserEmail(email string, state string, bookID int) (err error)
	GetTagsForBook (bookID int) (tags []Book, err error)
}