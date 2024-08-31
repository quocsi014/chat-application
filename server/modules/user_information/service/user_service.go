package service

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/quocsi014/common"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/user_information/entity"
)


type IUserRepository interface{
	InsertUser(ctx context.Context, user *entity.User) error
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUsersByUsername(ctx context.Context, username string, paging *common.Paging) ([]*entity.User, error)
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

func (service *UserService) CreateUser(ctx context.Context, user *entity.User) error {
	if _, err := service.repository.FindUserById(ctx, user.Id); err == nil {		
		return entity.ErrExistUser
	} else if !errors.Is(err, app_error.ErrRecordNotFound) {
		return app_error.ErrDatabase(err)
	}

	if user.Username == nil {
		return entity.ErrBlankUsername
	}
	if len(*user.Username) < 3 {
		return entity.ErrUsernameTooShort
	}
	if !isValidUsername(*user.Username) {
		return entity.ErrInvalidUsername
	}

	// Kiểm tra xem username đã tồn tại chưa
	if _, err := service.repository.GetUserByUsername(ctx, *user.Username); err == nil {
		return entity.ErrUsernameTaken
	} else if !errors.Is(err, app_error.ErrRecordNotFound) {
		return app_error.ErrDatabase(err)
	}

	// Kiểm tra các trường khác như cũ
	if user.Firstname == nil {
		return entity.ErrFirstNameMissing
	}
	if user.Lastname == nil {
		return entity.ErrLastnameMissing
	}
	if *user.Firstname == "" {
		return entity.ErrBlankFirstname
	}
	if *user.Lastname == "" {
		return entity.ErrBlankLastname
	}
	if !validNameRegex.MatchString(*user.Firstname) {
		return entity.ErrInvalidFirstname
	}
	if !validNameRegex.MatchString(*user.Lastname) {
		return entity.ErrInvalidLastname
	}

	err := service.repository.InsertUser(ctx, user)
	if err != nil {
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

	if user.Username != nil {
		if *user.Username == "" {
			return entity.ErrBlankUsername
		}
		if len(*user.Username) < 3 {
			return entity.ErrUsernameTooShort
		}
		if !isValidUsername(*user.Username) {
			return entity.ErrInvalidUsername
		}
		// Kiểm tra xem username mới có bị trùng không
		if *existingUser.Username != *user.Username {
			if _, err := service.repository.GetUserByUsername(ctx, *user.Username); err == nil {
				return entity.ErrUsernameTaken
			} else if !errors.Is(err, app_error.ErrRecordNotFound) {
				return app_error.ErrDatabase(err)
			}
		}
	}

	if user.Firstname != nil {
		if *user.Firstname == "" {
			return entity.ErrBlankFirstname
		}
		if !validNameRegex.MatchString(*user.Firstname) {
			return entity.ErrInvalidFirstname
		}
	}

	if user.Lastname != nil {
		if *user.Lastname == "" {
			return entity.ErrBlankLastname
		}
		if !validNameRegex.MatchString(*user.Lastname) {
			return entity.ErrInvalidLastname
		}
	}

	// Cập nhật user
	err = service.repository.UpdateUser(ctx, user)
	if err != nil {
		return app_error.ErrDatabase(err)
	}
	return nil
}

// Hàm kiểm tra tính hợp lệ của username
func isValidUsername(username string) bool {
	validUsernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return validUsernameRegex.MatchString(username)
}

func (service *UserService) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := service.repository.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return nil, entity.ErrUserNotFound
		}
		return nil, app_error.ErrDatabase(err)
	}
	return user, nil
}

func (service *UserService) GetUserById(ctx context.Context, userId string) (*entity.User, error) {
	user, err := service.repository.FindUserById(ctx, userId)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return nil, entity.ErrUserNotFound
		}
		return nil, app_error.ErrDatabase(err)
	}
	return user, nil
}

func (service *UserService) GetUsersByUsername(ctx context.Context, username string, paging *common.Paging) ([]*entity.User, error) {
	users, err := service.repository.GetUsersByUsername(ctx, username, paging)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return nil, app_error.ErrDatabase(err)
	}
	return users, nil
}
