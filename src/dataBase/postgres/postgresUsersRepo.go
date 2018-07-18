package postgres

import (
	"database/sql"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
)

type postgresUsersRepository struct {
	Db *sql.DB
}

//NewPostgresUsersRepo is a function to get New postgresUsersRepository which uses given connection
func NewPostgresUsersRepo(Db *sql.DB) repository.UserRepository {
	return &postgresUsersRepository{Db}
}

//GetUserByEmail gets users from users table by
func (p *postgresUsersRepository) GetUserByEmail(email string) (u repository.User, err error) {
	var n1, n2, n3 sql.NullInt64
	rows := p.Db.QueryRow(`SELECT nickname,password,email,notification_get_new_books,notification_get_when_book_reserved,notification_daily,
		has_book_for_exchange,returning_book_id,book_id,taken_book_id FROM gotoboox.users where email=$1`, email)

	err = rows.Scan(&u.Nickname, &u.Password, &u.Email, &u.NotificationGetBewBooks, &u.NotificationGetWhenBookReserved, &u.NotificationDaily, &u.HasBookForExchange, &n1, &n2, &n3)
	if err != nil {
		return
	}

	u.ReturningBookID = getNullableIntValue(n1)
	u.BookID = getNullableIntValue(n2)
	u.TakenBookID = getNullableIntValue(n3)
	return
}

//UpdateInsertUserByEmail updates a user or insert if there is no such user
func (p *postgresUsersRepository) UpdateUserByEmail(u repository.User, oldEmail string) (err error) {
	_, err = p.Db.Query(`UPDATE gotoboox.users set nickname=$1,email=$2,password=$3,notification_get_new_books=$4,
 								notification_get_when_book_reserved=$5,notification_daily=$6  where email=$7`,
		u.Nickname, u.Email, u.NewPassword, u.NotificationGetBewBooks, u.NotificationGetWhenBookReserved, u.NotificationDaily, oldEmail)
	return
}

//GetRoleByEmail get user role by email
func (p *postgresUsersRepository) GetRoleByEmail(email string) (user repository.User, err error) {
	row := p.Db.QueryRow("SELECT roles.title, users_roles.id FROM (gotoboox.users JOIN gotoboox.users_roles USING (id)) JOIN gotoboox.roles USING (role_id) WHERE users.email = $1", email)
	if err != nil {
		log.Printf("Get %v", err)

		return
	}
	err = row.Scan(&user.Role, &user.ID)
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	return
}

//InsertRolesToUsers insert to database user's role (default role is "user")
func (p *postgresUsersRepository) InsertRolesToUsers(userID int, roleID int) (err error) {

	tx, err := p.Db.Begin()
	if err != nil {
		return
	}

	stmt, err := tx.Prepare("INSERT INTO gotoboox.users_roles (id, role_id) values($1, $2)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userID, roleID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

//_, err = p.Db.Query("INSERT INTO gotoboox.users_roles (id, role_id) values($1, $2)",
//	userID, roleID)
//return
//}

//DeleteUserByEmail deletes user from database
func (p *postgresUsersRepository) DeleteUserByEmail(email string) (err error) {
	_, err = p.Db.Query("DELETE FROM gotoboox.users WHERE email=$1", email)
	return
}

//InsertUser is a function that inserts a user entity into a database
func (p *postgresUsersRepository) InsertUser(u repository.User) (lastID int, err error) {
	err = p.Db.QueryRow("INSERT INTO gotoboox.users (nickname,email,password,register_date) values($1,$2,$3,$4) RETURNING id",
		u.Nickname, u.Email, u.Password, u.RegisterDate).Scan(&lastID)
	return
}

//GetUsersEmailToNotifyNewBook is a function that
func (p *postgresUsersRepository) GetUsersEmailToNotifyNewBook() (u []repository.User, err error) {
	rows, err := p.Db.Query(
		"SELECT email,nickname FROM gotoboox.users where notification_get_new_books='true'")
	defer rows.Close()
	for rows.Next() {

		var user repository.User
		err = rows.Scan(&user.Email, &user.Nickname)
		if err != nil {
			return
		}
		u = append(u, user)
	}
	return
}

func (p *postgresUsersRepository) GetUsersEmailToNotifyReserved() (u []repository.User, err error) {
	rows, err := p.Db.Query(
		"SELECT email,nickname FROM gotoboox.users where notification_get_when_book_reserved='true'")
	defer rows.Close()
	for rows.Next() {
		var user repository.User
		err = rows.Scan(&user.Email, &user.Nickname)
		if err != nil {
			return
		}
		u = append(u, user)
	}
	return
}

func (p *postgresUsersRepository) SetUsersBookAsNullByBookID(id int) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.users set book_id=NULL where book_id=$1", id)
	return
}

func (p *postgresUsersRepository) GetAllUsers() (u []repository.User, err error) {
	rows, err := p.Db.Query(
		"SELECT id,email,nickname,exchanges_number FROM gotoboox.users LIMIT 2000")
	defer rows.Close()
	for rows.Next() {
		var user repository.User
		err = rows.Scan(&user.ID, &user.Email, &user.Nickname, &user.ExchangesNumber)
		if err != nil {
			return
		}
		u = append(u, user)
	}
	return
}

func (p *postgresUsersRepository) SetReturningBookIDByEmail(id int, email string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.users set returning_book_id=$1 where email=$2", id, email)
	return
}

func (p *postgresUsersRepository) ClearReturningBookIDByEmail(email string) (err error) {
	_, err = p.Db.Query("UPDATE gotoboox.users set returning_book_id=NULL where email=$1", email)
	return
}

func (p *postgresUsersRepository) MakeBookCross(email string) (err error) {
	u, err := p.GetUserByEmail(email)
	if err != nil {
		return
	}

	tx, err := p.Db.Begin()
	if err != nil {
		return
	}

	//First transaction
	{
		stmt, err := tx.Prepare(`UPDATE gotoboox.users SET  book_id=NULL,has_book_for_exchange=FALSE,
		  							taken_book_id=$1 where email=$2;`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(u.BookID, email); err != nil {
			tx.Rollback()
			return err
		}
	}
	//Second transaction
	{
		stmt, err := tx.Prepare(`UPDATE gotoboox.books SET  state=$1 where id=$2`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(repository.BookStateTaken, u.BookID); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func getNullableIntValue(nullableValue sql.NullInt64) int {
	if !nullableValue.Valid {
		return 0
	}
	return int(nullableValue.Int64)
}
