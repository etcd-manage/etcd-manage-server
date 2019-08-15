package models

import (
	"github.com/jinzhu/gorm"
)

// RolesModel 角色表
type RolesModel struct {
	Id        int32    `gorm:"column:id;primary_key" json:"id"`     //
	Name      string   `gorm:"column:name" json:"name"`             // 角色名
	CreatedAt JSONTime `gorm:"column:created_at" json:"created_at"` // 添加时间
	UpdatedAt JSONTime `gorm:"column:updated_at" json:"updated_at"` // 更新时间

}

// TableName 获取表名
func (RolesModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "roles")
}

// All 查询全部角色
func (m *RolesModel) All() (list []*RolesModel, err error) {
	err = client.Table(m.TableName()).Scan(&list).Error
	return
}

// Save 保存
func (m *RolesModel) Save() (err error) {
	err = client.Table(m.TableName()).Save(m).Error
	return
}

// Del 删除
func (m *RolesModel) Del(id int32) (err error) {
	err = client.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
	return
}
