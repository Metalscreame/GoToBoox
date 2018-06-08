package postgres

import (
	"database/sql"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

type postgresUsersRepository struct {
	Db *sql.DB
}

func NewPostgresUsersRepo(Db *sql.DB) repository.UserRepository {
	return &postgresUsersRepository{Db}
}

//GetUserByEmail gets users from users table by
func (p *postgresUsersRepository) GetUserByEmail(email string) (u repository.User, err error) {
	row := p.Db.QueryRow("SELECT id, nickname,password,email FROM gotoboox.users where email=$1", email)
	err = row.Scan(&u.ID, &u.Nickname, &u.Password, &u.Email)
	if err != nil {
		return
	}
	return
}

//UpdateInsertUserByEmail updates a user or insert if there is no such user
func (p *postgresUsersRepository) UpdateUserByEmail(u repository.User, oldEmail string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.users set nickname=$1,email=$2,password=$3 where email=$4",
		u.Nickname, u.Email, u.Password, oldEmail)
	return
}

//DeleteUserByEmail deletes user from database
func (p *postgresUsersRepository) DeleteUserByEmail(email string) (err error) {
	_, err = p.Db.Query("DELETE FROM gotoboox.users WHERE email=$1", email)
	return
}

//InsertUser is a function that inserts a user entity into a database
func (p *postgresUsersRepository) InsertUser(u repository.User) (err error) {
	_, err = p.Db.Query("INSERT INTO gotoboox.users (nickname,email,password,register_date) values($1,$2,$3,$4)",
		u.Nickname, u.Email, u.Password, u.RegisterDate)
	return
}
