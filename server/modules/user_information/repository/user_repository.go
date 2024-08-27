package repository

import (
	"context"

	"github.com/quocsi014/modules/user_information/entity"
	"gorm.io/gorm"
)


type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository)InsertUser(ctx context.Context, user *entity.User) error{
	return repo.db.Save(user).Error
}
