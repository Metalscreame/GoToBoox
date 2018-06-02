package users

import "github.com/metalscreame/GoToBoox/src/dataBase/repository/entity"



type UserRepository interface {
	GetUserByEmail(email string) (u entity.User, err error)
	UpdateUserByEmail(u entity.User) (err error)
	DeleteUserByEmail(email string) (err error)
	InsertUser(u entity.User)(err error)
}