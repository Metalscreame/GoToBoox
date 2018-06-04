package repository




type UserRepository interface {
	GetUserByEmail(email string) (u User, err error)
	UpdateUserByEmail(u User,email string) (err error)
	DeleteUserByEmail(email string) (err error)
	InsertUser(u User)(err error)
}