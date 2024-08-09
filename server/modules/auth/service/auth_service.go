package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
	"golang.org/x/crypto/bcrypt"
)

type IAuthRepository interface{
	GetAccount(ctx *gin.Context, email string) (*entity.Account, error)
}

type AuthService struct{
	repository IAuthRepository
	jwtSecretKey string
}

func NewAuthService(repo IAuthRepository, jwtSecretKey string) *AuthService{
	return &AuthService{
		repository: repo,
		jwtSecretKey: jwtSecretKey,
	}
}

func (as *AuthService)generateJwtToken(userId string) (string, error){
	jwtClaims := jwt.MapClaims{
		"user_id": userId,
		"exp": time.Now().Add(time.Hour*24).Unix(),
	}
	t := jwt.NewWithClaims( jwt.SigningMethodHS256, jwtClaims)
	return t.SignedString([]byte(as.jwtSecretKey))
}


func (as *AuthService) Login(ctx *gin.Context, account entity.Account) (string, error){
	a, err := as.repository.GetAccount(ctx, account.Email)
	if err != nil{
		if errors.Is(err, app_error.ErrRecordNotFound){
			return "",app_error.ErrUnauthenticatedError(err, "Email or password is incorrect")	
		}
		return "", app_error.ErrDatabase(err)
	}

	if err := bcrypt.CompareHashAndPassword( []byte(a.Password), []byte(account.Password)); err != nil{
		return "",app_error.ErrUnauthenticatedError(err, "Email or Password is incorrect")
	}

	jwtToken, err := as.generateJwtToken(a.Id)
	if err != nil{
		return "", app_error.ErrInternal(err)
	}

	return jwtToken, nil

}
