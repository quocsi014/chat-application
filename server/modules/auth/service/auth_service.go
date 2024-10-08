package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repository               IAuthRepository
	accountCachingRepository IAcccountCachingRepository
	jwtSecretKey             string
}

func (as *AuthService) GetJwtSecretKey() string {
	return as.jwtSecretKey
}

func NewAuthService(repo IAuthRepository, accountCachingRepository IAcccountCachingRepository, jwtSecretKey string) *AuthService {
	return &AuthService{
		repository:               repo,
		accountCachingRepository: accountCachingRepository,
		jwtSecretKey:             jwtSecretKey,
	}
}

func (as *AuthService) generateJwtToken(userId string) (string, error) {
	jwtClaims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return t.SignedString([]byte(as.jwtSecretKey))
}

func (as *AuthService) Login(ctx context.Context, account entity.LoginAccount) (string, error) {
	if account.Password == nil {
		return "", entity.ErrBlankPassword
	}

	if account.Account == nil {
		return "", app_error.ErrInvalidData(errors.New("Missing email"), "EMAIL_MISSING", "Email is required")
	}

	a, err := as.repository.GetAccount(ctx, *account.Account)
	if err != nil {
		if errors.Is(err, app_error.ErrRecordNotFound) {
			return "", app_error.ErrUnauthenticatedError(err, "Email or password is incorrect")
		}
		return "", app_error.ErrDatabase(err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*a.Password), []byte(*account.Password)); err != nil {
		return "", app_error.ErrUnauthenticatedError(err, "Email or Password is incorrect")
	}

	jwtToken, err := as.generateJwtToken(a.Id)
	if err != nil {
		return "", app_error.ErrInternal(err)
	}

	return jwtToken, nil
}

func (as *AuthService) isEmailTaken(ctx context.Context, email string) (bool, error) {
	if _, err := as.repository.GetAccount(ctx, email); err == nil {
		return true, app_error.ErrConflictData(nil, "EMAIL_EXIST", "Email has been taken")
	} else {
		if !errors.Is(err, app_error.ErrRecordNotFound) {
			return false, app_error.ErrDatabase(err)
		}
	}
	return false, nil
}

func (as *AuthService) generateJwtTokenWithEmail(email string) (string, error) {
	jwtClaims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 5).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	return t.SignedString([]byte(as.jwtSecretKey))
}

func validateEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	return re.MatchString(email)
}

func (as *AuthService) Register(ctx context.Context, account entity.Account) (string, error) {
	if account.Email == nil {
		return "", entity.ErrNilEmail
	}

	if account.Password == nil {
		return "", entity.ErrNilPassword
	}

	if len(*account.Password) < 6 {
		return "", entity.ErrInvalidPassword
	}

	if !validateEmail(*account.Email) {
		return "", entity.ErrInvaliEmail
	}

	if taken, err := as.isEmailTaken(ctx, *account.Email); taken || err != nil {
		return "", err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*account.Password), 10)
	if err != nil {
		return "", app_error.ErrInternal(err)
	}

	hashedPasswordString := string(hashedPassword)
	account.Password = &hashedPasswordString

	if err := as.accountCachingRepository.StoreAccount(ctx, account); err != nil {
		return "", app_error.ErrDatabase(err)
	}
	token, err := as.generateJwtTokenWithEmail(*account.Email)
	if err != nil {
		return "", app_error.ErrInternal(err)
	}
	return token, nil
}

func (as *AuthService) VerifyAccount(ctx context.Context, email string) (string, error) {
	account, err := as.accountCachingRepository.GetAccount(ctx, email)
	account.Id = uuid.NewString()
	if err != nil {
		return "", app_error.ErrDatabase(err)
	}
	if err := as.repository.InserAccount(ctx, account); err != nil {
		return "", app_error.ErrDatabase(err)
	}
	accessToken, err := as.generateJwtToken(account.Id)
	if err != nil {
		return "", app_error.ErrInternal(err)
	}
	return accessToken, nil
}
