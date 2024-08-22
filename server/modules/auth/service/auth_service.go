package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IAuthRepository interface{
	GetAccount(ctx context.Context, email string) (*entity.Account, error)
	InserAccount(ctx context.Context, account *entity.Account) error
}

type IOTPRepository interface{
	SetOtp(ctx context.Context, email, otp string) error
	GetOtp(ctx context.Context, email string) (string, error)
}

type AuthService struct{
	repository IAuthRepository
	otpRepository IOTPRepository
	jwtSecretKey string
}

func (as *AuthService)GetJwtSecretKey() string{
	return as.jwtSecretKey
}

func NewAuthService(repo IAuthRepository, otpRepo IOTPRepository, jwtSecretKey string) *AuthService{
	return &AuthService{
		repository: repo,
		otpRepository: otpRepo,
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

func (as *AuthService) Login(ctx context.Context, account entity.Account) (string, error){
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

func (as *AuthService)CreateEmailVerification(ctx context.Context, email, otp string) error {
	_, err := as.repository.GetAccount(ctx, email)
	if err == nil {
		return err
	}
	if errors.Is(err, app_error.ErrRecordNotFound){	
		return as.otpRepository.SetOtp(ctx, email, otp)
	}
	fmt.Println(err.Error())
	return app_error.ErrDatabase(err)
}

func (as *AuthService)generateVerifyOtpToken(email string) (string, error){

	jwtClaims := jwt.MapClaims{
		"email": email,
		"exp": time.Now().Add(time.Minute*5).Unix(),
	}
	t := jwt.NewWithClaims( jwt.SigningMethodHS256, jwtClaims)
	return t.SignedString([]byte(as.jwtSecretKey))
}

func (as *AuthService)VerifyOTP(ctx context.Context, email, otp string) (string, error) {
	storagedOtp, err := as.otpRepository.GetOtp(ctx, email)

	if err != nil {
		if err == redis.Nil{
			return "", app_error.ErrUnauthenticatedError(err, "OTP is incorrect or expired")
		}else{
			return "", app_error.ErrDatabase(err)
		}
	}

	if (otp != storagedOtp){
		return "", app_error.ErrUnauthenticatedError(nil, "OTP is incorrect or expired")
	}

	token, err := as.generateVerifyOtpToken(email)

	if err != nil{
		fmt.Println(err.Error())
		return "", app_error.ErrInternal(err)
	}

	return token, nil
}

func (as *AuthService)Register(ctx context.Context, account *entity.Account) error{
	accountID := uuid.NewString()
	account.Id = accountID
	if err := as.repository.InserAccount(ctx, account); err != nil{
		if errors.Is(err, gorm.ErrDuplicatedKey){
			return account.ErrUsernameExist()
		}else{
			fmt.Println(err.Error())
			return app_error.ErrDatabase(err)
		}
	}
	return nil

}
