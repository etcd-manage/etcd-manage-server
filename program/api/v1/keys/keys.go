package keys

import (
	"errors"
	"log"
	"net/http"

	"github.com/etcd-manage/etcd-manage-server/program/common"
	"github.com/gin-gonic/gin"
)

// KeysController key控制器
type KeysController struct {
}

// ReqKeyBody 添加和修改key请求body
type ReqKeyBody struct {
	Path  string `json:"path"`
	Value string `json:"value"`
	IsDir bool   `json:"is_dir"`
}

// List 获取目录下key列表
func (api *KeysController) List(c *gin.Context) {
	path := c.Query("path")
	log.Println(path)
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	list, err := client.List(path)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, list)
}

// Val 获取一个key的值
func (api *KeysController) Val(c *gin.Context) {
	path := c.Query("path")
	log.Println(path)
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	list, err := client.Val(path)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, list)
}

// Add 添加key
func (api *KeysController) Add(c *gin.Context) {
	req := new(ReqKeyBody)
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	err = c.Bind(req)
	if err != nil {
		return
	}
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	err = client.Add(req.Path, []byte(req.Value))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// Put 修改key
func (api *KeysController) Put(c *gin.Context) {
	req := new(ReqKeyBody)
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	err = c.Bind(req)
	if err != nil {
		return
	}
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	err = client.Put(req.Path, []byte(req.Value))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// Del 删除key
func (api *KeysController) Del(c *gin.Context) {
	path := c.Query("path")
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	if path == "" {
		err = errors.New("Path cannot be empty")
		return
	}
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	err = client.Del(path)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, "ok")
}

// Members 获取etcd服务节点
func (api *KeysController) Members(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	client, err := common.GetEtcdClientByGinContext(c)
	if err != nil {
		return
	}
	members, err := client.Members()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, members)
}
