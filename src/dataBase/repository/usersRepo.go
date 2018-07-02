package repository

//UserRepository is a repository interface that contains methods for member(user) that works with database
type UserRepository interface {
	GetUserByEmail(email string) (u User, err error)
	UpdateUserByEmail(u User,email string) (err error)
	DeleteUserByEmail(email string) (err error)
	InsertUser(u User)(err error)
	GetUsersEmailToNotifyNewBook()(u []User, err error)
	GetUsersEmailToNotifyReserved()(u []User, err error)
	SetUsersBookAsNullByBookId(id int)(err error)
	GetAllUsers()(u []User,err error)
	MakeBookCross(email string) (err error)
	SetReturningBookIDByEmail(id int,email string)(err error)
	ClearReturningBookIDByEmail(email string)(err error)
}

