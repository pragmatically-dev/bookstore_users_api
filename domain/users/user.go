package users

import (
	"strings"

	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

//User the core of this micro-service
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
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
