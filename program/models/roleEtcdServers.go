package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// RoleEtcdServersModel 角色权限表
type RoleEtcdServersModel struct {
	Id           int       `gorm:"column:id;primary_key" json:"id"`             //
	EtcdServerId int       `gorm:"column:etcd_server_id" json:"etcd_server_id"` // etcd服务id
	Type         int       `gorm:"column:type" json:"type"`                     // 0读 1写 2读写
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`         // 添加时间
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`         // 更新时间

}

// TableName 获取表名
func (RoleEtcdServersModel) TableName() string {
	return gorm.DefaultTableNameHandler(nil, "role_etcd_servers")
}
