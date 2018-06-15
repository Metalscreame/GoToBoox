package repository

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
	SetReturningBookIdByEmail(id int,email string)(err error)
	ClearReturningBookIdByEmail(email string)(err error)
}

