package errors

import "net/http"

//UserError implements error interface from builtin
type UserError struct {
	Reason string `json:"reason,omitempty"`
	Msg    string `json:"message,omitempty"`
	Code   int    `json:"code,omitempty"`
}

func (u *UserError) Error() string {
	return u.Msg
}

func NewBadRequestError(msg, reason string) *UserError {
	return &UserError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusBadRequest,
	}
}

func NewNotFoundError(msg, reason string) *UserError {
	return &UserError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusNotFound,
	}
}

func NewInternalServerError(msg, reason string) *UserError {
	return &UserError{
		Reason: reason,
		Msg:    msg,
		Code:   http.StatusInternalServerError,
	}
}
