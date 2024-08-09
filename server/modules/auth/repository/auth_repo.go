package repository

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common"
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

func (ar *AuthRepository)GetAccount(ctx *gin.Context, email string) (*entity.Account, error){
	account := entity.Account{}
	if err := ar.db.Where("email = ?", email).First(&account).Error; err != nil{
		fmt.Println(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, app_error.ErrRecordNotFound
		}else{
			return nil, err
		}
	}
	return &account, nil
}

