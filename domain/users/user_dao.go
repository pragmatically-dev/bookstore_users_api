package users

import (
	"github.com/go-sql-driver/mysql"
	db "github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
	"github.com/pragmatically-dev/bookstore_users_api/utils/dateutils"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

type mysqlErr struct {
	Number int
	Msg    string
}

func (err *mysqlErr) Error() string {
	return err.Msg
}

var (
	//ErrDuplicateEntry .
	ErrDuplicateEntry = mysqlErr{Number: 1062, Msg: "Error Duplicate Entry"}
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
	var errs errors.APIErrors
	var userFound User
	getErr := db.Client.Get(&userFound, getUserQuery, user.ID)
	if getErr != nil {

		errs.AddError(errors.NewNotFoundError("User not found", "Not Found"))
		return &errs
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
		sqlErr, ok := saveErr.(*mysql.MySQLError)
		if !ok {
			errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error: Error when trying to save user"))
			return &errs
		}
		switch sqlErr.Number {
		case uint16(ErrDuplicateEntry.Number):
			errs.AddError(errors.NewBadRequestError("The email is already registred", ErrDuplicateEntry.Error()))
			return &errs

		}
	}

	userID, err := res.LastInsertId()
	if err != nil {
		errs.AddError(errors.NewInternalServerError(err.Error(), "Internal Server Error"))
		return &errs
	}

	user.ID = userID
	return nil
}
