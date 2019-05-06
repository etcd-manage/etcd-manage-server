package keys

import (
	"log"
	"net/http"

	"github.com/etcd-manage/etcd-manage-server/program/common"
	"github.com/gin-gonic/gin"
)

// KeysController key控制器
type KeysController struct {
}

// List 获取目录下key列表
func (api *KeysController) List(c *gin.Context) {
	path := c.Query("path")
	log.Println(path)
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err,
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
