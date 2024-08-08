package repository

import (

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
	var account *entity.Account
	if err := ar.db.Where("email = ?", email).First(account).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, common.ErrRecordNotFound
		}else{
			return nil, err
		}
	}
	return account, nil
}

