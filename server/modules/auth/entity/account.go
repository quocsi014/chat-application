package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)


type Account struct{
	Id string `json:"id" gorm:"column:id"`
	Email string `json:"email" gorm:"column:email"`
	Username string `json:"username" gorm:"column:username"`
	Password string `json:"password" gorm:"column:password"`
}

func NewAccount(email, username, password string) *Account{
	return &Account{
		Email: email,
		Username: username,
		Password: password,
	}
}

func (a *Account)TableName() string{
	return "accounts"
}

func (a *Account)ErrUsernameExist() error{
	return app_error.ErrConflictData(errors.New("Username already exist"), "USERNAME_EXIST", "Username already exist")
}


