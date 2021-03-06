package userservice

import (
	"github.com/pragmatically-dev/bookstore_users_api/domain/users"
	"github.com/pragmatically-dev/bookstore_users_api/utils/cryptoutils"
	"github.com/pragmatically-dev/bookstore_users_api/utils/dateutils"
	"github.com/pragmatically-dev/bookstore_users_api/utils/errors"
)

type UserService struct {
}

func (service *UserService) GetUser(userID int64) (*users.User, *errors.APIErrors) {
	var errs errors.APIErrors
	if userID <= 0 {
		errs.AddError(errors.NewBadRequestError("Invalid user ID", "Could not parse ID"))
	}
	user := users.User{
		ID: userID,
	}
	getErrs := user.Get(true)
	if getErrs != nil {
		return nil, getErrs
	}
	return &user, nil

}

func (service *UserService) CreateUser(user users.User) (*users.User, *errors.APIErrors) {

	if errs := user.Validate(); errs != nil {
		return nil, errs
	}
	user.Status = users.StatusActive
	user.CreatedAt = dateutils.GetNowString()
	user.Password, _ = cryptoutils.HashPassword(user.Password)
	saveErrs := user.Save()
	if saveErrs != nil {
		return nil, saveErrs
	}

	return &user, nil
}

func (service *UserService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.APIErrors) {
	current, err := service.GetUser(user.ID) //respecting SOLID
	if err != nil {
		return nil, err
	}
	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (service *UserService) DeleteUser(userID int64) *errors.APIErrors {
	user, err := service.GetUser(userID) //respecting SOLID
	if err != nil {
		return err
	}
	errs := user.Delete()
	if errs != nil {
		return errs
	}
	return nil
}

func (service *UserService) FindByStatus(status string) (users.Users, *errors.APIErrors) {
	dao := &users.User{}
	return dao.FindByStatus(status)

}

func (service *UserService) LoginUser(req users.LoginRequest) (*users.User, *errors.APIErrors) {
	var errs errors.APIErrors
	dao := &users.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := dao.Get(false)
	if err != nil {
		return nil, err
	}

	match := cryptoutils.CheckPasswordHash(req.Password, dao.Password)
	if match {
		return dao, nil
	}
	errs.AddError(errors.NewBadRequestError("Bad Request", "Invalid credentials"))
	return nil, &errs
}
