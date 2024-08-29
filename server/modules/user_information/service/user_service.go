package service

import (
	"context"
	"errors"
	"regexp"

	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
)


type IUserRepository interface{
	InsertUser(ctx context.Context, user *entity.User) error
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
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
	if _, err := service.repository.FindUserById(ctx, user.Id); err == nil{		
		return entity.ErrExistUser
	}else{
		if !errors.Is(err, app_error.ErrRecordNotFound){
			return app_error.ErrDatabase(err)
		}
	}
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

func (service *UserService) UpdateUser(ctx context.Context, userId string, user *entity.User) error {
	existingUser, err := service.repository.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return entity.ErrUserNotFound
		}
		return app_error.ErrDatabase(err)
	}

	if user.Firstname != nil {
		if *user.Firstname == "" {
			return entity.ErrBlankFirstname
		}
		if !validNameRegex.MatchString(*user.Firstname) {
			return entity.ErrInvalidFirstname
		}
		existingUser.Firstname = user.Firstname
	}

	if user.Lastname != nil {
		if *user.Lastname == "" {
			return entity.ErrBlankLastname
		}
		if !validNameRegex.MatchString(*user.Lastname) {
			return entity.ErrInvalidLastname
		}
		existingUser.Lastname = user.Lastname
	}

	if user.AvatarURL != nil {
		existingUser.AvatarURL = user.AvatarURL
	}

	err = service.repository.UpdateUser(ctx, existingUser)
	if err != nil {
		return app_error.ErrDatabase(err)
	}
	return nil
}
