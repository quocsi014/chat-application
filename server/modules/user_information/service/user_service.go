package service

import (
	"context"
	"regexp"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
)


type IUserRepository interface{
	InsertUser(ctx context.Context, user *entity.User) error
}
type UserService struct{
	repository IUserRepository
}

func NewUserService(repository IUserRepository) *UserService{
	return &UserService{
		repository: repository,
	}
}

var validNameRegex = regexp.MustCompile(`^[\p{L}\p{M}\p{Zs}'\-\.]+$`)
func (service *UserService)CreateUser(ctx context.Context, user *entity.User) error{
	if user.Firstname == nil{
		return entity.ErrFirstNameMissing
	}
	if user.Lastname == nil{
		return entity.ErrLastnameMissing
	}
	if *user.Firstname == ""{
		return entity.ErrBlankFirstname
	}
	if *user.Lastname == ""{
		return entity.ErrBlankLastname
	}
	if !validNameRegex.MatchString(*user.Firstname){
		return entity.ErrInvalidFirstname
	}
	if !validNameRegex.MatchString(*user.Lastname){
		return entity.ErrInvalidLastname
	}

	err := service.repository.InsertUser(ctx, user)
	if err != nil{
		return app_error.ErrDatabase(err)
	}
	return nil
}
