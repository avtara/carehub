package models

import "errors"

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserDuplicate     = errors.New("email has been taken")
	ErrorUserWrongPassword = errors.New("wrong password or email")
	ErrorCategoryNotFound  = errors.New("category not found")
	ErrorComplainNotFound  = errors.New("complain not found")
)

const (
	GroupFilePhotoUser = "photo_user"
	MaxSize            = 1024 * 1024
)
