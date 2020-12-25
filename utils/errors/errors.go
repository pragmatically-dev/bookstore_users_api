package errors

import "net/http"

func NewBadRequestError(msg, reason string) *UserError {
	return &UserError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusBadRequest,
	}
}
