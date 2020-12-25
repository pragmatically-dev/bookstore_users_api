package utils

//APIErrors contains the whole logic for handle APIErrors
type APIErrors struct {
	Errors []error `json:"errors,omitempty"`
	Code   int     `json:"code,omitempty"`
}

//AddError allows add news to the errors slice
func (ae *APIErrors) AddError(err error, code int) {
	if err != nil {
		ae.Errors = append(ae.Errors, err)
	}
	ae.addCode(code)
}

//AddCode setter for add the status code of the error
func (ae *APIErrors) addCode(code int) {
	ae.Code = code
}

//GetErrors returns the slice of errors
func (ae *APIErrors) GetErrors() []error {

	return ae.Errors
}

type BaseValidationError interface {
	Error() string
}

//UserError implements error interface from builtin
type UserError struct {
	Reason string `json:"reason,omitempty"`
	Msg    string `json:"message,omitempty"`
}

func (u *UserError) Error() string {
	return u.Msg
}
