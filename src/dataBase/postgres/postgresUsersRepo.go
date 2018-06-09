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
	row := p.Db.QueryRow(
		"SELECT id, nickname,password,email,notification_get_new_books,notification_get_when_book_reserved FROM gotoboox.users where email=$1", email)
	err = row.Scan(&u.ID, &u.Nickname, &u.Password, &u.Email, &u.NotificationGetBewBooks, &u.NotificationGetWhenBookReserved)
	if err != nil {
		return
	}
	return
}

//UpdateInsertUserByEmail updates a user or insert if there is no such user
func (p *postgresUsersRepository) UpdateUserByEmail(u repository.User, oldEmail string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.users set nickname=$1,email=$2,password=$3,notification_get_new_books=$4, notification_get_when_book_reserved=$5  where email=$6",
		u.Nickname, u.Email, u.Password, u.NotificationGetBewBooks, u.NotificationGetWhenBookReserved, oldEmail)
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

//func (p *postgresUsersRepository)GetUsersBookByEmail(email string)(ub repository.UsersBooks,err error){
//	//_, err = p.Db.Query("SELECT b.title, b.description, b.popularity, b.title FROM gotoboox.users a, gotoboox.books b where a.categories_id=b.id",
//		//id SERIAL PRIMARY KEY,
//		//title CHARACTER VARYING (250) NOT NULL,
//		//description CHARACTER VARYING (900) NOT NULL,
//		//popularity REAL NOT NULL DEFAULT 0,
//		//isTaken BOOLEAN DEFAULT FALSE,
//		//file_path CHARACTER VARYING (250) NOT NULL,
//		//categories_id INT REFERENCES gotoboox.categories (id)
//	return
//}

func (p *postgresUsersRepository) GetUsersEmailToNotifyNewBook() (u []repository.User, err error) {
	rows,err := p.Db.Query(
		"SELECT email,nickname FROM gotoboox.users where notification_get_new_books='true'")
	i:=0
	for rows.Next(){
		err = rows.Scan(&u[i].Email,&u[i].Nickname)
		if err != nil {
			return
		}
		i++
	}
	return
}

func (p *postgresUsersRepository) GetUsersEmailToNotifyReserved() (u []repository.User, err error) {
	rows,err := p.Db.Query(
		"SELECT email,nickname FROM gotoboox.users where notification_get_when_book_reserved='true'")
	i:=0
	for rows.Next(){
		err = rows.Scan(&u[i].Email,&u[i].Nickname)
		if err != nil {
			return
		}
		i++
	}
	return
}