package errors

type CustomError struct {
	error
	Cause   error
	Message string
	Code    int
}

func BuildCustomError(err error, code int, message string) CustomError {
	return CustomError{
		Cause:   err,
		Code:    code,
		Message: message,
	}
}

func (e *CustomError) WithRootCause(err error) *CustomError {
	e.Cause = err
	return e
}
