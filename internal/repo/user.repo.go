package repo

import (
	"go/go-backend-api/global"
	"go/go-backend-api/internal/model"

	"go.uber.org/zap"
)

type IUserRepository interface {
	EmailHasExist(email string) bool
}

type userRepository struct {
}

// EmailHasExist implements IUserRepository.
func (us *userRepository) EmailHasExist(email string) bool {
	rs := global.MyDB.Table(model.TableNameGoCrmUser).Where("usr_email = ?", email).Find(&model.GoCrmUser{})
	if rs.Error != nil {
		global.Logger.Error("EmailHasExist error", zap.Error(rs.Error))
		return false
	}
	if rs.RowsAffected <= 0 {
		return false
	}
	return true
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
