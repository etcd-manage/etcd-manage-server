package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	ReadOnly = iota
	ReadWrite
)

// RoleEtcdServersModel 角色权限表
type RoleEtcdServersModel struct {
	Id           int32     `gorm:"column:id;primary_key" json:"id"`             //
	EtcdServerId int32     `gorm:"column:etcd_server_id" json:"etcd_server_id"` // etcd服务id
	Type         int32     `gorm:"column:type" json:"type"`                     // 0读 1写
	RoleId       int32     `gorm:"column:role_id" json:"role_id"`               // 角色id
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`         // 添加时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`         // 更新时间

}

// TableName 获取表名
func (RoleEtcdServersModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "role_etcd_servers")
}

// FirstByRoleIdAndEtcdServerIdAndType 根据role_id、etcd_server_id和type查询一条数据
func (m *RoleEtcdServersModel) FirstByRoleIdAndEtcdServerIdAndType(roleId, etcdServerId, typ int32) (err error) {
	err = client.Table(m.TableName()).Where("role_id = ? and etcd_server_id = ? and type >= ?", roleId, etcdServerId, typ).First(m).Error
	return
}
