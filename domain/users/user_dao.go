package users

import (
	db "github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
	"github.com/pragmatically-dev/bookstore_users_api/utils/dateutils"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
	"github.com/pragmatically-dev/bookstore_users_api/utils/mysqlutils"
)

var (
	createUserQuery = `INSERT INTO users(
	first_name,
	last_name,
	email,
	created_at) VALUES(?,?,?,?)`

	getUserQuery = `SELECT
		user_id,
		first_name,
		last_name,
		email,
		created_at FROM users WHERE  user_id = ?;`
)

//Get returns an user from dbs
func (user *User) Get() *errors.APIErrors {

	var userFound User
	getErr := db.Client.Get(&userFound, getUserQuery, user.ID)
	if getErr != nil {
		return mysqlutils.ParseError(getErr)
	}
	user.CopyWith(&userFound)

	return nil
}

//Save insert an user in the db
func (user *User) Save() *errors.APIErrors {
	var errs errors.APIErrors
	user.CreatedAt = dateutils.GetNowString()
	stmt, err := db.Client.Prepare(createUserQuery)
	defer stmt.Close()
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	res, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt)
	if saveErr != nil {
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}

	user.ID = userID
	return nil
}
