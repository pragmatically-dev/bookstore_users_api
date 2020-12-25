package errors

//APIErrors contains the whole logic for handle APIErrors
type APIErrors struct {
	Errors []error `json:"errors,omitempty"`
}

//AddError allows add news to the errors slice
func (ae *APIErrors) AddError(err error) {
	if err != nil {
		ae.Errors = append(ae.Errors, err)
	}

}

//GetErrors returns the slice of errors
func (ae *APIErrors) GetErrors() []error {

	return ae.Errors
}
