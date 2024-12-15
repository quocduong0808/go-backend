package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleName string `gorm:"column:role_name; type:varchar(255);"`
	RoleNote bool   `gorm:"column:role_note; type:text;"`
}

func (u *Role) TableName() string {
	return "go_db_role"
}
