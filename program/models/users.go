package models

import (
	"fmt"

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

// UsersJoinRoleModel 用户列表带角色名
type UsersJoinRoleModel struct {
	RoleName string `gorm:"column:role_name" json:"role_name"`
	UsersModel
}

// FirstByUsernameAndPassword 根据用户名密码查询数据
func (m *UsersModel) FirstByUsernameAndPassword(username, password string) (err error) {
	err = client.Table(m.TableName()).Where("username = ? and password = ?", username, password).First(m).Error
	return
}

// List 分页列表
func (m *UsersModel) List(userId, roleId int32, name string, offset, limit int) (list []*UsersModel, err error) {
	where, params := m.listWhere(userId, roleId, name, offset, limit)
	err = client.Table(m.TableName()+" as u").Select("u.*, r.name as role_name").
		Joins(fmt.Sprintf("left join %s as r on u.role_id = r.id", new(RolesModel).TableName())).
		Where(where, params...).
		Offset(offset).
		Limit(limit).
		Scan(&list).Error
	return
}

// ListCount 分页总数据量
func (m *UsersModel) ListCount(userId, roleId int32, name string, offset, limit int) (_c int32, err error) {
	where, params := m.listWhere(userId, roleId, name, offset, limit)
	err = client.Table(m.TableName()).Where(where, params...).Count(&_c).Error
	return
}

// 分页查询条件组织
func (m *UsersModel) listWhere(userId, roleId int32, name string, offset, limit int) (where string, params []interface{}) {
	where = ""
	params = make([]interface{}, 0)
	if userId > 0 {
		where = "id = ? "
		params = append(params, userId)
	}
	if roleId > 0 {
		if where != "" {
			where += " and "
		}
		where += " role_id = ? "
		params = append(params, roleId)
	}
	if name != "" {
		if where != "" {
			where += " and "
		}
		where += " username like ?"
		params = append(params, "%"+name+"%")
	}
	return
}

// Save 保存
func (m *UsersModel) Save() (err error) {
	err = client.Table(m.TableName()).Save(m).Error
	return
}

// Del 删除
func (m *UsersModel) Del(id int32) (err error) {
	err = client.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
	return
}
