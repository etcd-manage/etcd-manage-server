package models

import (
	"encoding/json"
	"fmt"
)

// EtcdServersModel etcd 服务列表
type EtcdServersModel struct {
	ID        int32    `json:"id" gorm:"column:id;primary_key"`
	Version   string   `json:"version" gorm:"column:version"`
	Name      string   `json:"name" gorm:"column:name"`
	Address   string   `json:"address" gorm:"column:address"`
	TlsEnable string   `json:"tls_enable" gorm:"column:tls_enable"`
	CertFile  string   `json:"cert_file" gorm:"column:cert_file"`
	KeyFile   string   `json:"key_file" gorm:"column:key_file"`
	CaFile    string   `json:"ca_file" gorm:"column:ca_file"`
	Username  string   `json:"username" gorm:"column:username"`
	Password  string   `json:"password" gorm:"column:password"`
	Desc      string   `json:"desc" gorm:"column:desc"`
	CreatedAt JSONTime `gorm:"column:created_at" json:"created_at"` // 添加时间
	UpdatedAt JSONTime `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

// TableName 表名
func (EtcdServersModel) TableName() string {
	return "etcd_servers"
}

// All 获取全部
func (m *EtcdServersModel) All(name string, roleId int32) (list []*EtcdServersModel, err error) {
	re := new(RoleEtcdServersModel).TableName()
	err = client.Table(m.TableName()+" as e").Select("e.*").
		Joins("LEFT JOIN "+re+" as re on re.etcd_server_id = e.id").
		Where("e.name like ? and re.role_id = ?", fmt.Sprintf("%%%s%%", name), roleId).
		Scan(&list).Error
	return
}

// FirstById 根据id查询一个etcd服务
func (m *EtcdServersModel) FirstById(id int32) (one *EtcdServersModel, err error) {
	one = new(EtcdServersModel)
	err = client.Model(m).Where("id = ?", id).First(one).Error
	return
}

// Insert 添加
func (m *EtcdServersModel) Insert() (err error) {
	err = client.Create(m).Error
	return
}

// Update 修改
func (m *EtcdServersModel) Update() (err error) {
	edit := make(map[string]interface{}, 0)
	js, _ := json.Marshal(m)
	json.Unmarshal(js, &edit)
	err = client.Model(new(EtcdServersModel)).Where("id = ?", m.ID).Updates(edit).Error
	return
}

// Del 删除
func (m *EtcdServersModel) Del(id int32) (err error) {
	err = client.Table(m.TableName()).Where("id = ?", id).Delete(m).Error
	return
}
