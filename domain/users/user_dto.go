package users

import (
	"strings"

	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

//User the core of this micro-service
type User struct {
	ID        int64  `json:"user_id" db:"user_id" `
	FirstName string `json:"first_name" db:"first_name" `
	LastName  string `json:"last_name" db:"last_name" `
	Email     string `json:"email" db:"email" `
	CreatedAt string `json:"created_at" db:"created_at" `
}

func (user *User) CopyWith(data *User) {
	user.ID = data.ID
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.CreatedAt = data.CreatedAt
}

func (user *User) Validate() *errors.APIErrors {
	var errs errors.APIErrors
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//TODO: IMPROVE EMAIL VALIDATION
	if user.Email == "" {
		errs.AddError(errors.NewBadRequestError("Invalid email", "The Email must be completed"))
	}

	if len(errs.Errors) > 0 {
		return &errs
	}
	return nil
}
