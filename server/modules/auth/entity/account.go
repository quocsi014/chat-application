package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)


type Account struct{
	Id string `json:"id" gorm:"column:id"`
	Email *string `json:"email" gorm:"column:email"`
	Password *string `json:"password" gorm:"column:password"`
}

func NewAccount(email, password string) *Account{
	return &Account{
		Email: &email,
		Password: &password,
	}
}

func (a *Account)TableName() string{
	return "accounts"
}
var (
	ErrEmailExist = app_error.ErrConflictData(errors.New("Email already exist"), "EMAIL_EXIST", "Email have been taken")

	ErrInvaliEmail = app_error.ErrInvalidData(errors.New("Invalid email"), "INVALID_EMAIL", "This is not an email")

	ErrNilEmail = app_error.ErrInvalidData(errors.New("Missing email"), "EMAIL_MISSING", "Email is required")

	ErrBlankPassword = app_error.ErrInvalidData(errors.New("Password blank"), "BLANK_PASSWORD", "Password can not be blank")

	ErrInvalidPassword = app_error.ErrInvalidData(errors.New("Invalid password"), "INVALID_PASSWORD", "Password must have at least 6 characters")

	ErrNilPassword = app_error.ErrInvalidData(errors.New("Missing Password"), "PASSWOR_MISSING", "Password is required")
)

type LoginAccount struct{
	Account *string `json:"account"`
	Password *string `json:"password"`
}
