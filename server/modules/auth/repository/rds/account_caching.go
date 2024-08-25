package rds

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/quocsi014/common/app_error"
	"github.com/quocsi014/modules/auth/entity"
)

type AccountCaching struct{
	rdb *redis.Client
	exp time.Duration
}

func NewAccountCaching(rdb *redis.Client, exp time.Duration) *AccountCaching{
	return &AccountCaching{
		rdb: rdb,
		exp: exp,
	}
}

func (ac *AccountCaching)StoreAccount(ctx context.Context, account entity.Account) error{
	key := "account:" + *account.Email
	err := ac.rdb.HSet(ctx, key, map[string]interface{}{
		"email": *account.Email,
		"username": *account.Username,
		"password": *account.Password,
	}).Err()
	return err
}

func (ac *AccountCaching)GetAccount(ctx context.Context, email string) (*entity.Account, error){
	key := "account:" + email
	exist, _ := ac.rdb.Exists(ctx, key).Result()
	if exist != 1{
		return nil, app_error.ErrUnauthenticatedError(nil, "Key is not exist")
	}
	mapData := ac.rdb.HGetAll(ctx, key)
	vals := mapData.Val()
	return entity.NewAccount(vals["email"], vals["username"], vals["password"]), nil
}
