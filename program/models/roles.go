package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// RolesModel 角色表
type RolesModel struct {
	Id        int       `gorm:"column:id;primary_key" json:"id"`     //
	Name      string    `gorm:"column:name" json:"name"`             // 角色名
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 添加时间
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"` // 更新时间

}

// TableName 获取表名
func (RolesModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "roles")
}
