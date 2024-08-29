package repository

import (
	"context"
	"errors"

	"github.com/quocsi014/common/app_error"
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

func (repo *UserRepository)FindUserById(ctx context.Context, id string) (*entity.User, error){
	user := entity.User{}
	if err := repo.db.Where("id = ?", id).First(&user).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, app_error.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return repo.db.Model(&entity.User{}).Where("id = ?", user.Id).Updates(user).Error
}

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := entity.User{}
	if err := repo.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app_error.ErrRecordNotFound
		}
		return nil, err
	}
	return &user, nil
}
