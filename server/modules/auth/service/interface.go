package service

import (
	"context"
	"github.com/quocsi014/modules/auth/entity"
)

type IAuthRepository interface {
	GetAccount(ctx context.Context, email string) (*entity.Account, error)
	InserAccount(ctx context.Context, account *entity.Account) error
}

type IOTPRepository interface {
	SetOtp(ctx context.Context, email, otp string) error
	GetOtp(ctx context.Context, email string) (string, error)
}

type IAcccountCachingRepository interface {
	StoreAccount(ctx context.Context, account entity.Account) error
	GetAccount(ctx context.Context, email string) (*entity.Account, error)
}

type IAccountService interface {
	Login(ctx context.Context, account entity.LoginAccount) (string, error)
	GetJwtSecretKey() string
	Register(ctx context.Context, account entity.Account) (string, error)
	VerifyAccount(ctx context.Context, email string) (string, error)
}
