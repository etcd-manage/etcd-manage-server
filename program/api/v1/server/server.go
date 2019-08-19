package server

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/etcd-manage/etcd-manage-server/program/models"
	"github.com/etcd-manage/etcdsdk/etcdv3"
	"github.com/etcd-manage/etcdsdk/model"
	"github.com/gin-gonic/gin"
)

// ServerController etcd服务列表相关操作
type ServerController struct {
}

// List 获取etcd服务列表，全部
func (api *ServerController) List(c *gin.Context) {
	name := c.Query("name")
	list, err := new(models.EtcdServersModel).All(name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusOK, list)
}

// Add 添加服务
func (api *ServerController) Add(c *gin.Context) {
	var err error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	// 添加
	req := new(models.EtcdServersModel)
	err = c.Bind(req)
	if err != nil {
		return
	}
	err = req.Insert()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// Update 修改服务
func (api *ServerController) Update(c *gin.Context) {
	var err error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}
	// 添加
	req := new(models.EtcdServersModel)
	err = c.Bind(req)
	if err != nil {
		return
	}
	err = req.Update()
	if err != nil {
		return
	}
	// TODO 删除已存在的此服务连接

	c.JSON(http.StatusOK, "ok")
}

// Restore 修复v1版本或e3w对目录的标记
func (api *ServerController) Restore(c *gin.Context) {
	etcdId := c.Query("etcd_id")
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	etcdIdNum, _ := strconv.Atoi(etcdId)
	etcdOne := new(models.EtcdServersModel)
	etcdOne, err = etcdOne.FirstById(int32(etcdIdNum))
	if err != nil {
		return
	}
	if etcdOne.Version != model.ETCD_VERSION_V3 {
		err = errors.New("Only V3 version is allowed to be repaired")
		return
	}
	// 连接etcd
	cfg := &model.Config{
		Version:   etcdOne.Version,
		Address:   strings.Split(etcdOne.Address, ","),
		TlsEnable: etcdOne.TlsEnable == "true",
		CertFile:  etcdOne.CaFile,
		KeyFile:   etcdOne.KeyFile,
		CaFile:    etcdOne.CaFile,
		Username:  etcdOne.Username,
		Password:  etcdOne.Password,
	}
	client, err := etcdv3.NewClient(cfg)
	if err != nil {
		return
	}
	clientV3, ok := client.(*etcdv3.EtcdV3Sdk)
	if ok == false {
		err = errors.New("Connecting etcd V3 service error")
		return
	}
	err = clientV3.Restore()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// SetRolesReq 设置角色请求参数
type SetRolesReq struct {
	EtcdId int32           `json:"etcd_id"`
	Roles  map[int32]int32 `json:"roles` // 下标角色id值为 0只读或1读写
}

// SetRoles 设置etcd服务角色
func (api *ServerController) SetRoles(c *gin.Context) {
	req := new(SetRolesReq)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	m := new(models.RoleEtcdServersModel)
	err = m.UpByEtcdId(req.EtcdId, req.Roles)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
