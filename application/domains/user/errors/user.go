package errors

import appErrors "github.com/rof20004/healthy-api/application/errors"

var (
	ErrInvalidId       = appErrors.BuildCustomError(nil, 400, "Invalid user id")
	ErrNameIsRequired  = appErrors.BuildCustomError(nil, 400, "User name is required")
	ErrInvalidAge      = appErrors.BuildCustomError(nil, 400, "User age needs to be greater than 0")
	ErrInvalidEmail    = appErrors.BuildCustomError(nil, 400, "Invalid user email")
	ErrEmailIsRequired = appErrors.BuildCustomError(nil, 400, "User email is required")
	ErrUsersNotFound   = appErrors.BuildCustomError(nil, 404, "Users not found")
)
