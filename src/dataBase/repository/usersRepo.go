package repository

type UserRepository interface {
	GetRoleByEmail(email string) (user User, err error)
	GetUserByEmail(email string) (u User, err error)
	UpdateUserByEmail(u User, email string) (err error)
	DeleteUserByEmail(email string) (err error)
	InsertUser(u User) (lastID int, err error)
	GetUsersEmailToNotifyNewBook() (u []User, err error)
	GetUsersEmailToNotifyReserved() (u []User, err error)
	SetUsersBookAsNullByBookId(id int) (err error)
	GetAllUsers() (u []User, err error)
	MakeBookCross(email string) (err error)
	SetReturningBookIdByEmail(id int, email string) (err error)
	ClearReturningBookIdByEmail(email string) (err error)
	InsertRolesToUsers (userID int, roleID int) (err error)
}
