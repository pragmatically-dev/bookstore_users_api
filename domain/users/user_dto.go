package users

import (
	"strings"

	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

const (
	StatusActive = "active"
)

//User the core of this micro-service
type User struct {
	ID        int64  `json:"user_id" db:"user_id" `
	FirstName string `json:"first_name" db:"first_name" `
	LastName  string `json:"last_name" db:"last_name" `
	Email     string `json:"email" db:"email" `
	CreatedAt string `json:"created_at" db:"created_at" `
	Status    string `json:"status" db:"status" `
	Password  string `json:"password" db:"password" `
}

func (user *User) CopyWith(data *User) {
	user.ID = data.ID
	user.FirstName = data.FirstName
	user.LastName = data.LastName
	user.Email = data.Email
	user.CreatedAt = data.CreatedAt
	user.Status = data.Status
	user.Password = data.Password
}

func (user *User) Validate() *errors.APIErrors {
	var errs errors.APIErrors
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	//TODO: IMPROVE EMAIL VALIDATION
	if user.Email == "" {
		errs.AddError(errors.NewBadRequestError("Invalid email", "The Email must be completed"))
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		errs.AddError(errors.NewBadRequestError("Invalid Password", "The password cannot be empty"))
	}

	if len(errs.Errors) > 0 {
		return &errs
	}
	return nil
}
