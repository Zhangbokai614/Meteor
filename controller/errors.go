package controller

import "errors"

var (
	StatusUserError = 497
)

var (
	ErrorUserAlreadyExists = errors.New("User alread exists")
	ErrorUserDoesNotExists = errors.New("User does not exist")
)
