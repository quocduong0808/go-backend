package service

import (
	"go/go-backend-api/global"
	"go/go-backend-api/internal/repo"
	"go/go-backend-api/pkg/response"
	"go/go-backend-api/pkg/utils"

	"go.uber.org/zap"
)

type IUserService interface {
	Register(email, phone, username, password string) int
}

type userService struct {
	userRepo    repo.IUserRepository
	authRepo    repo.IAuthRepository
	mailService IMailService
}

// Register implements IUserService.
func (u *userService) Register(email string, phone string, username string, password string) int {
	//0. hash email
	emailHash := utils.Hash(email)
	//5. check otp available
	//6. user spam check
	//1. check email exist
	if u.userRepo.EmailHasExist(email) {
		global.Logger.Error("EmailHasExist error", zap.String("email", email))
		return response.ErrorCodeEmailExist
	}
	//2. new otp
	otp, err := utils.GenerateRandomOTP()
	if err != nil {
		global.Logger.Error("GenerateRandomOTP error", zap.Error(err))
		return response.ErrCodeGenerateOTP
	}
	//3. save otp in redis with expire time
	err = u.authRepo.AddOTP(emailHash, otp, 300)
	if err != nil {
		global.Logger.Error("AddOTP error", zap.Error(err))
		return response.ErrCodeSaveOTPRedis
	}
	//4. send email
	err = u.mailService.SendTextMail([]string{email}, "OTP", otp)
	if err != nil {
		global.Logger.Error("SendTextMail error", zap.Error(err))
		return response.ErrCodeSendMail

	}
	return response.SuccessCode
}

func NewUserService(userRepo repo.IUserRepository, authRepo repo.IAuthRepository, mailService IMailService) IUserService {
	return &userService{
		userRepo:    userRepo,
		authRepo:    authRepo,
		mailService: mailService,
	}
}
