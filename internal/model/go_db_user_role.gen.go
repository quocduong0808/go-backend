// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameGoDbUserRole = "go_db_user_role"

// GoDbUserRole mapped from table <go_db_user_role>
type GoDbUserRole struct {
	UserID int64 `gorm:"column:user_id;primaryKey" json:"user_id"`
	RoleID int64 `gorm:"column:role_id;primaryKey" json:"role_id"`
}

// TableName GoDbUserRole's table name
func (*GoDbUserRole) TableName() string {
	return TableNameGoDbUserRole
}
