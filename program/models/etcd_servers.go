package models

// EtcdServersModel etcd 服务列表
type EtcdServersModel struct {
	ID        int32  `json:"id,omitempty" gorm:"column:id;primary_key"`
	Version   string `json:"version,omitempty" gorm:"column:version"`
	Name      string `json:"name,omitempty" gorm:"column:name"`
	Address   string `json:"address,omitempty" gorm:"column:address"`
	TlsEnable string `json:"tls_enable,omitempty" gorm:"column:tls_enable"`
	CertFile  string `json:"cert_file,omitempty" gorm:"column:cert_file"`
	KeyFile   string `json:"key_file,omitempty" gorm:"column:key_file"`
	CaFile    string `json:"ca_file,omitempty" gorm:"column:ca_file"`
	Username  string `json:"username,omitempty" gorm:"column:username"`
	Password  string `json:"password,omitempty" gorm:"column:password"`
	Desc      string `json:"desc,omitempty" gorm:"column:desc"`
}

// TableName 表名
func (EtcdServersModel) TableName() string {
	return "etcd_servers"
}

// All 获取全部
func (m *EtcdServersModel) All() (list []*EtcdServersModel, err error) {
	err = client.Model(m).Scan(&list).Error
	return
}

// FirstById 根据id查询一个etcd服务
func (m *EtcdServersModel) FirstById(id int32) (one *EtcdServersModel, err error) {
	one = new(EtcdServersModel)
	err = client.Model(m).Where("id = ?", id).First(one).Error
	return
}
