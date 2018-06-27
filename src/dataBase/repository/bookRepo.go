package repository

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
	UpdateBookState(bookId int, state string) (err error)
	UpdateBookStateAndUsersBookIdByUserEmail(email string, state string, bookId int) (err error)
	GetTagsForBook (bookID int) (tags []Book, err error)
}