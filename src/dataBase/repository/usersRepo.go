package repository

import (
	"github.com/metalscreame/GoToBoox/src/dataBase/entity"
	db "github.com/metalscreame/GoToBoox/src/dataBase/configuration"
	"bytes"
	"errors"
	"database/sql"
)

const (
	selectQueryType = "select"
	insertQueryType = "insert"
	deleteQueryType = "delete"
	updateQueryType = "update"
)

type User entity.User

//GetUserByEmail gets users from users table by
func (u *User) GetUserByEmail() (err error) {
	query := prepareQueryString(selectQueryType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	rows, err := execQueueByEmail(stmt, u)
	if err != nil {
		return
	}

	for rows.Next() {
		err = rows.Scan(&u.Id, &u.Nickname, &u.Email, &u.Password, &u.RegistrDate) ///email,username,password,firstname,lastname,Created
		if err != nil {
			return
		}
	}
	return
}

func (u *User) UpdateUserByEmail() (err error) {
	query := prepareQueryString(updateQueryType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	err= execInsertStmtByEmail(stmt,u)
	if err != nil{
		return
	}
	return
}

func (u *User) DeleteUserByEmail() (err error) {
	query := prepareQueryString(deleteQueryType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	err= execInsertStmtByEmail(stmt,u)
	if err != nil{
		return
	}
	return
}

func (u *User) CreateUser() (err error) {
	query := prepareQueryString(insertQueryType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	res, err := stmt.Exec(u.Nickname, u.Email, u.Password, u.RegistrDate)
	if err != nil {
		return
	}

	_, err = res.RowsAffected()
	if err != nil {
		return errors.New("There is no user with such email")
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
		b.WriteString(" set nickname=&1,email=$2,password=$3 where email=$4")
		return b.String()
	case selectQueryType:
		b.WriteString("select from ")
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

func execInsertStmtByEmail(stmt *sql.Stmt, u *User) (err error) {
	res, err := stmt.Exec(u.Nickname,u.Email,u.Password,u.Email)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil || affect == 0 {
		return errors.New("There is no user with such email")
	}
}

func execQueueByEmail(stmt *sql.Stmt, u *User) (rows *sql.Rows, err error) {
	rows, err = stmt.Query(u.Email)
	if err != nil {
		return nil, err
	}
	return
}