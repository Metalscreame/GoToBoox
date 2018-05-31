package repository

import (
	"github.com/metalscreame/GoToBoox/src/dataBase/entity"
	db "github.com/metalscreame/GoToBoox/src/dataBase/dbConnection"
	"bytes"
	"errors"
	"database/sql"
)

const (
	selectType = "select"
	insertType = "insert"
	deleteType = "delete"
	updateType = "update"
)

type User entity.User

//GetUser gets users from users table by
func (u *User) GetUser() (err error) {
	query := prepareQueryString(selectType)
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

func (u *User) UpdateUser() (err error) {
	query := prepareQueryString(updateType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	err=execStmtByEmail(stmt,u)
	if err != nil{
		return
	}
	return
}

func (u *User) DeleteUser() (err error) {
	query := prepareQueryString(deleteType)
	stmt, err := db.GlobalDataBaseConnection.Prepare(query)
	if err != nil {
		return
	}

	err=execStmtByEmail(stmt,u)
	if err != nil{
		return
	}
	return
}

func (u *User) CreateUser() (err error) {
	query := prepareQueryString(insertType)
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

func prepareQueryString(typeOfQueue string) (string) {
	var b bytes.Buffer

	switch typeOfQueue {
	case insertType:
		b.WriteString("insert into ")
		b.WriteString(db.DB_SCHEMA)
		b.WriteString(db.DB_USERS_TABLE)
		b.WriteString("(nickname,email,password,registrDate) values($1,$2,$3,$4)")
		return b.String()
	case updateType:
		b.WriteString("update ")
		b.WriteString(db.DB_SCHEMA)
		b.WriteString(db.DB_USERS_TABLE)
		b.WriteString(" set nickname=&1,email=$2,password=$3 where email=$4")
		return b.String()
	case selectType:
		b.WriteString("select from ")
	case deleteType:
		b.WriteString("delete from ")
	default:
		return ""
	}

	b.WriteString(db.DB_SCHEMA)
	b.WriteString(db.DB_USERS_TABLE)
	b.WriteString(" where email=$1")
	return b.String()
}

func execStmtByEmail(stmt *sql.Stmt, u *User) (err error) {
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
