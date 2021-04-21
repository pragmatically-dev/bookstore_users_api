package mysqlutils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
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
	//ErrNoRows
	ErrNoRows = mysqlErr{Number: 0000, Msg: "no rows in result set"}
)

func ParseError(err error) *errors.APIErrors {
	var globalErrors errors.APIErrors
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrNoRows.Msg) {
			globalErrors.AddError(errors.NewNotFoundError("Not Found", "No record matching given id"))
			return &globalErrors
		}
		globalErrors.AddError(errors.NewInternalServerError("Internal Server Error", "Error parsing database response"))
		return &globalErrors
	}
	switch sqlErr.Number {
	case uint16(ErrDuplicateEntry.Number):
		globalErrors.AddError(errors.NewBadRequestError("Internal Server Error", "Invalid data"))
	}
	globalErrors.AddError(errors.NewInternalServerError("Internal Server Error", "Error parsing database response"))
	return &globalErrors
}
