package service

import (
	"github.com/gin-gonic/gin"
	"github.com/quocsi014/common"
	"github.com/quocsi014/modules/auth/entity"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface{
	GetAccount(ctx *gin.Context, email string) (*entity.Account, error)
}

type AuthService struct{
	repository IAuthRepository
}

func NewAuthService(repo IAuthRepository) *AuthService{
	return &AuthService{
		repository: repo,
	}
}


func (as *AuthService) Login(ctx *gin.Context, account *entity.Account) error{
	a, err := as.repository.GetAccount(ctx, account.Email)
	if err != nil{
		if err == common.ErrRecordNotFound{
			return common.NewUnauthenticatedError(err, "Email or password is incorrect")	
		}
			return common.NewUnauthenticatedError(err, "Email or password is incorrect")	
	}

	if err := bcrypt.CompareHashAndPassword( []byte(a.Password), []byte(account.Password)); err != nil{
		return common.NewUnauthenticatedError(err, "Email or Password is incorrect")
	}

	account.Id = a.Id

	return nil

}
