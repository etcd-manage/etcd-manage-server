package models

import (
	"github.com/jinzhu/gorm"
)

// UsersModel 用户表
type UsersModel struct {
	Id        int32    `gorm:"column:id;primary_key" json:"id"`     //
	Username  string   `gorm:"column:username" json:"username"`     // 用户名
	Password  string   `gorm:"column:password" json:"password"`     // 密码
	Email     string   `gorm:"column:email" json:"email"`           // 邮箱
	RoleId    int32    `gorm:"column:role_id" json:"role_id"`       // 角色id
	CreatedAt JSONTime `gorm:"column:created_at" json:"created_at"` // 添加时间
	UpdatedAt JSONTime `gorm:"column:updated_at" json:"updated_at"` // 更新时间
	Token     string   `gorm:"-" json:"token"`                      // 登录token
}

// TableName 获取表名
func (UsersModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "users")
}

// FirstByUsernameAndPassword 根据用户名密码查询数据
func (m *UsersModel) FirstByUsernameAndPassword(username, password string) (err error) {
	err = client.Table(m.TableName()).Where("username = ? and password = ?", username, password).First(m).Error
	return
}
