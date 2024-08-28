package entity

import (
	"errors"

	"github.com/quocsi014/common/app_error"
)

type User struct{
	Id string `json:"id,omitempty" gorm:"column:id"`
	Firstname *string `json:"firstname,omitempty" gorm:"column:firstname"`
	Lastname *string `json:"lastname,omitempty" gorm:"column:lastname"`
}

func (u *User)TableName() string{
	return "users"
}

func NewUser(firstname, lastname string) *User{
	return &User{
		Firstname: &firstname,
		Lastname: &lastname,
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
)
