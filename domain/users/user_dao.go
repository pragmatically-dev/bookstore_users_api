package users

import (
	"context"
	"time"

	db "github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
	"github.com/pragmatically-dev/bookstore_users_api/logger"
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
		logger.Error("Error when trying to get an user from db", getErr)
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
		logger.Error("Error when trying to prepare save user statement", err)
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	res, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedAt, user.Password, user.Status)
	if saveErr != nil {
		logger.Error("Error when trying to save user", saveErr)
		return mysqlutils.ParseError(saveErr)
	}

	userID, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error when trying to get the lastInsertId", err)

		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}

	user.ID = userID

	return nil
}

//Update
func (user *User) Update() *errors.APIErrors {
	var errs errors.APIErrors
	stmt, err := db.Client.Prepare(updateUserQuery)
	if err != nil {
		logger.Error("Error when trying to prepare the update statement", err)
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		logger.Error("Error when trying to update user", err)
		return mysqlutils.ParseError(err)
	}
	return nil
}

//Delete
func (user *User) Delete() *errors.APIErrors {
	var errs errors.APIErrors
	stmt, err := db.Client.Prepare(deleteUserQuery)
	if err != nil {
		logger.Error("Error when trying to prepare the delete statement", err)
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID)
	if err != nil {
		logger.Error("Error when trying to delete user", err)

		return mysqlutils.ParseError(err)
	}
	return nil
}

//findByStatus .
func (user *User) FindByStatus(status string) ([]User, *errors.APIErrors) {
	var users []User
	var errs errors.APIErrors
	//TODO: this statement might cause troubles in the future
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := db.Client.SelectContext(ctx, &users, findUserByStatus, status)
	if err != nil {
		logger.Error("Error when trying to get users by status", err)

		return nil, mysqlutils.ParseError(err)
	}
	if len(users) <= 0 {
		errs.AddError(errors.NewNotFoundError("Not Found", "No users matching status"))
		return nil, &errs
	}
	return users, nil
}
