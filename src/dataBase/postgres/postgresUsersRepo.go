package postgres

import (
	"bytes"
	"database/sql"
	db "github.com/metalscreame/GoToBoox/src/dataBase"
	"time"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
)

const (
	selectQueryType = "select"
	insertQueryType = "insert"
	deleteQueryType = "delete"
	updateQueryType = "update"
)

type postgresUsersRepository struct {
	Db *sql.DB
}

func NewPostgresUsersRepo(Db *sql.DB) repository.UserRepository {
	return &postgresUsersRepository{Db}
}

//GetUserByEmail gets users from users table by
func (p *postgresUsersRepository) GetUserByEmail(email string) (u repository.User, err error) {

	query := prepareQueryString(selectQueryType)
	stmt, err := p.Db.Prepare(query)
	if err != nil {
		return
	}

	//rows, err := execQueueByEmail(stmt, email)
	row := stmt.QueryRow(email)
	if err != nil {
		return
	}

	err = row.Scan(&u.ID, &u.Nickname, &u.Email, &u.Password, &u.RegisterDate)
	if err != nil {
		return
	}
	return
}

//UpdateInsertUserByEmail updates a user or insert if there is no such user
func (p *postgresUsersRepository) UpdateUserByEmail(u repository.User, oldEmail string) (err error) {
	query := prepareQueryString(updateQueryType)
	stmt, err := p.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Exec(u.Nickname, u.Email, u.Password, oldEmail)

	//err = execInsertStmtByEmail(stmt, &u)
	if err != nil {
		return
	}
	return
}

//DeleteUserByEmail deletes user from database
func (p *postgresUsersRepository) DeleteUserByEmail(email string) (err error) {
	query := prepareQueryString(deleteQueryType)
	stmt, err := p.Db.Prepare(query)
	if err != nil {
		return
	}

	_, err = stmt.Query(email)
	if err != nil {
		return
	}
	return
}

func (p *postgresUsersRepository) InsertUser(u repository.User) (err error) {
	query := prepareQueryString(insertQueryType)
	stmt, err := p.Db.Prepare(query)
	if err != nil {
		return
	}
	err = execInsertStmtByEmail(stmt, &u)
	if err != nil {
		return
	}
	return
}

func prepareQueryString(typeOfQuery string) (string) {
	var b bytes.Buffer

	switch typeOfQuery {
	case insertQueryType:
		b.WriteString("insert into ")
		b.WriteString(db.DB_SCHEMA)
		b.WriteString(db.DB_USERS_TABLE)
		b.WriteString("(nickname,email,password,registrDate) values($1,$2,$3,$4)")
		return b.String()
	case updateQueryType:
		b.WriteString("update ")
		b.WriteString(db.DB_SCHEMA)
		b.WriteString(db.DB_USERS_TABLE)
		b.WriteString(" set nickname=$1,email=$2,password=$3 where email=$4")
		return b.String()
	case selectQueryType:
		b.WriteString("select * from ")
	case deleteQueryType:
		b.WriteString("delete from ")
	default:
		return ""
	}

	b.WriteString(db.DB_SCHEMA)
	b.WriteString(db.DB_USERS_TABLE)
	b.WriteString(" where email=$1")
	return b.String()
}

func execInsertStmtByEmail(stmt *sql.Stmt, u *repository.User) (err error) {
	//err = convertRegUserTime(u)
	//if err != nil {
	//	return
	//}
	res, err := stmt.Exec(u.Nickname, u.Email, u.Password, u.RegisterDate)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return
	}
	return
}

func convertRegUserTime(u *repository.User) (err error) {
	layout := "2006-01-02"
	updatedAt, err := time.Parse(layout, u.RegTimeStr)
	if err != nil {
		return
	}
	u.RegisterDate = updatedAt
	return
}
