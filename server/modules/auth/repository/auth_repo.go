package repository

import (
	"errors"
	"context"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
	"gorm.io/gorm"
)

type AuthRepository struct{
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository{
	return &AuthRepository{
		db: db,
	}
}

func (ar *AuthRepository)GetAccount(ctx context.Context, user string) (*entity.Account, error){
	account := entity.Account{}
	if err := ar.db.Where("email = ? or username = ?", user, user).First(&account).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, app_error.ErrRecordNotFound
		}else{
			return nil, err
		}
	}
	return &account, nil
}

func (ar *AuthRepository)InserAccount(ctx context.Context, account *entity.Account) error{
	return ar.db.Create(&account).Error
}

func (ar *AuthRepository)GetAccountByUsername(ctx context.Context, username string) (*entity.Account, error){
	account := entity.Account{}
	if err := ar.db.Where("username = ?", username).First(&account).Error; err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, app_error.ErrRecordNotFound
		}else{
			return nil, err
		}
	}
	return &account, nil
}
