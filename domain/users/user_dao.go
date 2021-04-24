package users

import (
	"context"
	"time"

	"github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
	db "github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
	"github.com/pragmatically-dev/bookstore_users_api/utils/mysqlutils"
)

var (
	createUserQuery = `INSERT INTO users(
	first_name,
	last_name,
	email,
	created_at,
	password,
	status
	) VALUES(?,?,?,?,?,?)`

	getUserQuery = `SELECT
		user_id,
		first_name,
		last_name,
		email,
		created_at,
		status
		FROM users WHERE  user_id = ?;`

	updateUserQuery = `UPDATE users SET 
			first_name=?, 
			last_name =?, 
			email=? 
		WHERE user_id=?`

	deleteUserQuery = ` DELETE FROM users_db.users WHERE user_id =?;`

	findUserByStatus = `SELECT
	user_id,
	first_name,
	last_name,
	email,
	status,
	created_at FROM users WHERE  status = ?;`
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
	stmt, err := db.Client.Prepare(createUserQuery)
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	res, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Password, user.Status)
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

//Update
func (user *User) Update() *errors.APIErrors {
	var errs errors.APIErrors
	stmt, err := usersdb.Client.Prepare(updateUserQuery)
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

//Delete
func (user *User) Delete() *errors.APIErrors {
	var errs errors.APIErrors
	stmt, err := usersdb.Client.Prepare(deleteUserQuery)
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

//findByStatus .
func (user *User) FindByStatus(status string) ([]User, *errors.APIErrors) {
	var users []User
	var errs errors.APIErrors
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	err := db.Client.SelectContext(ctx, &users, findUserByStatus, status)
	if err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	if len(users) <= 0 {
		errs.AddError(errors.NewNotFoundError("Not Found", "No users matching status"))
		return nil, &errs
	}
	return users, nil
}
