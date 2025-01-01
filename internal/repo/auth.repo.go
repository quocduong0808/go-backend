package repo

import (
	"context"
	"fmt"
	"go/go-backend-api/global"
	"time"
)

type IAuthRepository interface {
	AddOTP(email string, otp string, expireTime int) error
}

type authRepository struct{}

// AddOTP implements IAuthRepository.
func (a *authRepository) AddOTP(email string, otp string, expireTime int) error {
	ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
	return global.Redis.SetEx(ctx, fmt.Sprintf("email:%s:otp", email), otp, time.Duration(expireTime)*time.Second).Err()
}

func NewAuthRepository() IAuthRepository {
	return &authRepository{}
}
