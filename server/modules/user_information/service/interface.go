package service

import (
	"context"
	"github.com/quocsi014/common"
	"github.com/quocsi014/modules/user_information/entity"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *entity.User) error
	FindUserById(ctx context.Context, id string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	GetUsersByUsername(ctx context.Context, username string, paging *common.Paging) ([]entity.User, error)
}

type IUserService interface {
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, userId string, user *entity.User) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUserById(ctx context.Context, userId string) (*entity.User, error)
	GetUsersByUsername(ctx context.Context, username string, paging *common.Paging) ([]entity.User, error)
}
