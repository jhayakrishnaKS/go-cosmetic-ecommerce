package constants

import "errors"

var (
	ErrTitleTaken         = errors.New("title Taken")
	ErrRoleTaken          = errors.New("role Taken")
	ErrUsernameTaken      = errors.New("username Taken")
	ErrEmailTaken         = errors.New("email Taken")
	ErrStatusTaken        = errors.New("order status Taken already")
	ErrInvalidCredentials = errors.New("check your Credentials")
	ErrAccessTokenExpired = errors.New("access token is expired")
	ErrAddressEmpty       = errors.New("address fields cannot be empty")
	ErrAddressNotFound    = errors.New("address not found")
	ErrItemNotFound=errors.New("item not found")
	ErrUnableToUpdateStatus=errors.New("unable to update status")
	ErrOrderNotFound=errors.New("order not found")
	ErrProductDescriptionEmpty=errors.New("product title, description, and brand cannot be empty")
)
