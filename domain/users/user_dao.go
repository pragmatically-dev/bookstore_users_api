package users

import (
	"github.com/pragmatically-dev/bookstore_users_api/utils/dateutils"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
	db "github.com/pragmatically-dev/bookstore_users_api/datasources/mysql/usersdb"
)

var (
	usersDB = make(map[int64]*User)
)

//DATA ACCESS OBJECT
func (user *User) Get() *errors.APIErrors {
	if err:= db.Client.Ping();err!=nil{
		panic(err)
	}
	var errs errors.APIErrors
	res := usersDB[user.ID]
	if res == nil {
		errs.AddError(errors.NewNotFoundError( "The user has not been found","User not found"))
		return &errs
	}
	user.CopyWith(res)
	return nil
}

func (user *User) Save() *errors.APIErrors {
	var errs errors.APIErrors
	current := usersDB[user.ID]
	if current != nil {
		if current.Email == user.Email {
			errs.AddError(errors.NewBadRequestError("Email already exist", "The email already exist in the dbs"))
			return &errs
		}
		errs.AddError(errors.NewBadRequestError("User already exist", "The User already exist in the dbs"))
		return &errs
	}
	user.CreatedAt= dateutils.GetNowString()
	usersDB[user.ID]=user

	return nil
}
