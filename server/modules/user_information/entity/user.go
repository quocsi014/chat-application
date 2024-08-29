package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)

type User struct{
	Id string `json:"id,omitempty" gorm:"column:id"`
	Username *string `json:"username,omitempty" gorm:"column:username"`
	Firstname *string `json:"firstname,omitempty" gorm:"column:firstname"`
	Lastname *string `json:"lastname,omitempty" gorm:"column:lastname"`
	AvatarURL *string `json:"avatar_url,omitempty" gorm:"column:avatar_url"`
}

func (u *User)TableName() string{
	return "users"
}

func NewUser(username, firstname, lastname, avatar_url string ) *User{
	return &User{
		Username: &username,
		Firstname: &firstname,
		Lastname: &lastname,
		AvatarURL: &avatar_url,
	}
}

var(
	ErrBlankFirstname = app_error.ErrInvalidData(errors.New("Firstname is blank"), "BLANK_FIRSTNAME", "Firstname can not be blank")

	ErrInvalidFirstname = app_error.ErrInvalidData(errors.New("Firstname have invalid character"), "INVALID_FIRSTNAME", "Firstname have invalid character")

	ErrFirstNameMissing = app_error.ErrInvalidData(errors.New("Firstname missing"), "FIRTNAME_MISSING", "Firstname is required")

	ErrBlankLastname = app_error.ErrInvalidData(errors.New("Lasename is blank"), "BLANK_LASTNAME", "Lastname can not be blank")

	ErrInvalidLastname = app_error.ErrInvalidData(errors.New("Lastname have invalid character"), "INVALID_LASTNAME", "Lastname have invalid character")

	ErrLastnameMissing = app_error.ErrInvalidData(errors.New("Lastname missing"), "LASTNAME_MISSING", "Lastname is required")

	ErrExistUser = app_error.ErrConflictData(errors.New("User has been initialized"), "INITIALIZED", "A user for this account has been previously created.")

	ErrUserNotFound = app_error.ErrNotFound(errors.New("User not found"), "USER_NOT_FOUND", "The user does not exist")

	// Thêm các lỗi liên quan đến username
	ErrBlankUsername = app_error.ErrInvalidData(errors.New("Username is blank"), "BLANK_USERNAME", "Username cannot be blank")
	ErrInvalidUsername = app_error.ErrInvalidData(errors.New("Username has invalid characters"), "INVALID_USERNAME", "Username can only contain letters, numbers, and underscores")
	ErrUsernameTooShort = app_error.ErrInvalidData(errors.New("Username is too short"), "USERNAME_TOO_SHORT", "Username must be at least 3 characters long")
	ErrUsernameTaken = app_error.ErrConflictData(errors.New("Username is already taken"), "USERNAME_TAKEN", "This username is already in use")
)
