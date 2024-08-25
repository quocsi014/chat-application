package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)


type Account struct{
	Id string `json:"id" gorm:"column:id"`
	Email *string `json:"email" gorm:"column:email"`
	Username *string `json:"username" gorm:"column:username"`
	Password *string `json:"password" gorm:"column:password"`
}

func NewAccount(email, username, password string) *Account{
	return &Account{
		Email: &email,
		Username: &username,
		Password: &password,
	}
}

func (a *Account)TableName() string{
	return "accounts"
}
var (
	ErrUsernameExist = app_error.ErrConflictData(errors.New("Username already exist"), "USERNAME_EXIST", "Username have been taken")

	ErrEmailExist = app_error.ErrConflictData(errors.New("Email already exist"), "EMAIL_EXIST", "Email have been taken")

	ErrInvaliEmail = app_error.ErrInvalidData(errors.New("Invalid email"), "This is not an email")

	ErrNilEmail = app_error.ErrInvalidData(errors.New("Missing email"), "Email is required")

	ErrBlankUsername = app_error.ErrInvalidData(errors.New("Username is blank"), "Username can not be blank")

	ErrInvalidUsername = app_error.ErrInvalidData(errors.New("Invalid username"), "Username must have at least 5 characters")

	ErrNilUsername = app_error.ErrInvalidData(errors.New("Missing username"), "Username is required")

	ErrBlankPassword = app_error.ErrInvalidData(errors.New("Password blank"), "Password can not be blank")

	ErrInvalidPassword = app_error.ErrInvalidData(errors.New("Invalid password"), "Password must have at least 6 characters")

	ErrNilPassword = app_error.ErrInvalidData(errors.New("Missing Password"), "Password is required")
)

type LoginAccount struct{
	Account *string `json:"account"`
	Password *string `json:"password"`
}
