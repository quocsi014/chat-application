package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)

type UserInformation struct {
	Username  *string `json:"username,omitempty" gorm:"column:username"`
	Firstname *string `json:"firstname,omitempty" gorm:"column:firstname"`
	Lastname  *string `json:"lastname,omitempty" gorm:"column:lastname"`
	AvatarURL *string `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
}

type User struct {
	Id string `json:"id,omitempty" gorm:"column:id;primaryKey"`
	UserInformation
}

func (u *User) TableName() string {
	return "users"
}

func NewUser(username, firstname, lastname, avatar_url string) *User {
	return &User{
		UserInformation: UserInformation{
			Username:  &username,
			Firstname: &firstname,
			Lastname:  &lastname,
			AvatarURL: &avatar_url,
		},
	}
}

// Các định nghĩa lỗi đã được cập nhật
var (
	ErrBlankFirstname = app_error.ErrInvalidData(errors.New("firstname is blank"), "BLANK_FIRSTNAME", "firstname cannot be blank")

	ErrInvalidFirstname = app_error.ErrInvalidData(errors.New("firstname has invalid character"), "INVALID_FIRSTNAME", "firstname contains invalid characters")

	ErrFirstNameMissing = app_error.ErrInvalidData(errors.New("firstname is missing"), "FIRSTNAME_MISSING", "firstname is required")

	ErrBlankLastname = app_error.ErrInvalidData(errors.New("lastname is blank"), "BLANK_LASTNAME", "lastname cannot be blank")

	ErrInvalidLastname = app_error.ErrInvalidData(errors.New("lastname has invalid character"), "INVALID_LASTNAME", "lastname contains invalid characters")

	ErrLastnameMissing = app_error.ErrInvalidData(errors.New("lastname is missing"), "LASTNAME_MISSING", "lastname is required")

	ErrExistUser = app_error.ErrConflictData(errors.New("user has been initialized"), "INITIALIZED", "a user for this account has been previously created")

	ErrUserNotFound = app_error.ErrNotFound(errors.New("user not found"), "USER_NOT_FOUND", "the user does not exist")

	ErrBlankUsername = app_error.ErrInvalidData(errors.New("username is blank"), "BLANK_USERNAME", "username cannot be blank")
	ErrInvalidUsername = app_error.ErrInvalidData(errors.New("username has invalid characters"), "INVALID_USERNAME", "username can only contain letters, numbers, and underscores")
	ErrUsernameTooShort = app_error.ErrInvalidData(errors.New("username is too short"), "USERNAME_TOO_SHORT", "username must be at least 3 characters long")
	ErrUsernameTaken = app_error.ErrConflictData(errors.New("username is already taken"), "USERNAME_TAKEN", "this username is already in use")
)
