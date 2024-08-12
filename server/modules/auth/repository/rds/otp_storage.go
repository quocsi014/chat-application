package rds

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type OTPRepository struct{
	rdb *redis.Client
	otpExpiration time.Duration
}

func NewOTPRepository(rdb *redis.Client, expiration time.Duration) *OTPRepository{
	return &OTPRepository{
		rdb: rdb,
		otpExpiration: expiration,
	}
}

func (repo *OTPRepository)SetOtp(ctx context.Context, email, otp string) error{
	return repo.rdb.Set(ctx, email, otp, repo.otpExpiration).Err()
}

func (repo *OTPRepository)GetOtp(ctx context.Context, email string) (string, error){
	return repo.rdb.Get(ctx, email).Result()
}



