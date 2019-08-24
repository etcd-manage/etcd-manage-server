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
func (m *RoleEtcdServersModel) UpByEtcdId(list []*AllByEtcdIdData) (err error) {
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
		typ := -1
		if v.Write == 1 {
			typ = 1
		} else if v.Read == 1 {
			typ = 0
		}
		one := &RoleEtcdServersModel{
			EtcdServerId: v.EtcdServerId,
			Type:         int32(typ),
			RoleId:       v.RoleId,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		err = tx.Create(one).Error
		if err != nil {
			return
		}
	}
	return
}

type AllByEtcdIdData struct {
	EtcdServerId int32  `gorm:"column:etcd_server_id" json:"etcd_server_id"` // etcd服务id
	RoleId       int32  `gorm:"column:role_id" json:"role_id"`               // 角色id
	Name         string `gorm:"column:name" json:"name"`
	Type         int32  `gorm:"column:type" json:"type"`
	Read         int32  `gorm:"column:read" json:"read"`
	Write        int32  `gorm:"column:write" json:"write"`
}

// AllByEtcdId 查询etcd服务权限配置列表
func (m *RoleEtcdServersModel) AllByEtcdId(etcdId int32) (list []*AllByEtcdIdData, err error) {
	err = client.Table(m.TableName()+" as re").Select("r.name, re.type, re.etcd_server_id, re.role_id").
		Joins(" join "+new(RolesModel).TableName()+" as r on r.id = re.role_id").
		Where("etcd_server_id = ?", etcdId).
		Scan(&list).Error
	if err == nil {
		for _, v := range list {
			if v.Type == 1 {
				v.Write = 1
				v.Read = 1
			} else if v.Type == 0 {
				v.Write = 0
				v.Read = 1
			} else {
				v.Write = 0
				v.Read = 0
			}
		}
	}
	rList, _err := new(RolesModel).All()
	if _err != nil {
		err = _err
		return
	}
	if len(rList) > len(list) {
		for _, v := range rList {
			exist := false
			for _, v1 := range list {
				if v1.RoleId == v.Id {
					exist = true
				}
			}
			if exist == false {
				list = append(list, &AllByEtcdIdData{
					EtcdServerId: etcdId,
					RoleId:       v.Id,
					Name:         v.Name,
					Type:         -1,
					Read:         0,
					Write:        0,
				})
			}
		}
	}
	return
}

// Save 保存角色信息
func (m *RoleEtcdServersModel) Save() (err error) {
	err = client.Table(m.TableName()).Create(m).Error
	return
}

// DelByEtcdId 删除
func (m *RoleEtcdServersModel) DelByEtcdId(etcdId int32) (err error) {
	err = client.Table(m.TableName()).Where("etcd_server_id = ?", etcdId).Delete(m).Error
	return
}
