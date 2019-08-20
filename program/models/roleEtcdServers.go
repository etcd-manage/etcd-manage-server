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
	Id           int32    `gorm:"column:id;primary_key" json:"id"`             //
	EtcdServerId int32    `gorm:"column:etcd_server_id" json:"etcd_server_id"` // etcd服务id
	Type         int32    `gorm:"column:type" json:"type"`                     // 0读 1写
	RoleId       int32    `gorm:"column:role_id" json:"role_id"`               // 角色id
	CreatedAt    JSONTime `gorm:"column:created_at" json:"created_at"`         // 添加时间
	UpdatedAt    JSONTime `gorm:"column:updated_at" json:"updated_at"`         // 更新时间

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

// UpByEtcdId 根据etcd_id更新角色
func (m *RoleEtcdServersModel) UpByEtcdId(list []*RoleEtcdServersModel) (err error) {
	tx := client.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	// 先删除再插入
	err = tx.Table(m.TableName()).Where("etcd_server_id = ?", list[0].EtcdServerId).Delete(m).Error
	if err != nil {
		return
	}
	now := JSONTime(time.Now())
	for _, v := range list {
		v.UpdatedAt = now
		v.CreatedAt = now
		err = tx.Create(v).Error
		if err != nil {
			return
		}
	}
	return
}

type AllByEtcdIdData struct {
	Name  string `gorm:"column:name" json:"name"`
	Type  int32  `gorm:"column:type" json:"type"`
	Read  int32  `gorm:"column:read" json:"read"`
	Write int32  `gorm:"column:write" json:"write"`
}

// AllByEtcdId 查询etcd服务权限配置列表
func (m *RoleEtcdServersModel) AllByEtcdId(etcdId int32) (list []*AllByEtcdIdData, err error) {
	err = client.Table(m.TableName()+" as re").Select("r.name, re.type").
		Joins(" join "+new(RolesModel).TableName()+" as r on r.id = re.role_id").
		Where("etcd_server_id = ?", etcdId).
		Scan(&list).Error
	if err == nil {
		for _, v := range list {
			if v.Type == 1 {
				v.Write = 1
				v.Read = 1
			} else {
				v.Write = 0
				v.Read = 1
			}
		}
	}
	return
}
