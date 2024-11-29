package constants

import "errors"

var (
	// QUERY PARAMS
	ErrInvalidStartDateParam error = errors.New("invalid 'start_date' format. must be in 'yyyy-mm-dd' format")
	ErrInvalidEndDateParam   error = errors.New("invalid 'end_date' format. must be in 'yyyy-mm-dd' format")

	// VALIDATOR
	ErrFieldRequired error = errors.New("field is required")
	ErrInvalidFormat error = errors.New("format is invalid")

	// AUTH
	ErrInvalidToken      error = errors.New("token is invalid")
	ErrIncorrectEmail    error = errors.New("email is incorrect")
	ErrIncorrectPassword error = errors.New("password is incorrect")
	ErrUserNotFound      error = errors.New("user not found")
	ErrAccessNotAllowed  error = errors.New("access not allowed")

	// HASH PASSWORD
	ErrInvalidHash              error = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleArgonVersion error = errors.New("incompatible version of argon2")
)
