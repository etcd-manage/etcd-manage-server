package v1

import (
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/keys"
	gin "github.com/gin-gonic/gin"
)

// APIV1 v1版接口
type APIV1 struct {
}

// Register 注册路由
func (v1 *APIV1) Register(router *gin.RouterGroup) {
	// etcd key 管理
	gx := router.Group("/keys")
	keysController := new(keys.KeysController)
	gx.GET("", keysController.List)
	gx.POST("", keysController.Add)
	gx.PUT("", keysController.Put)
	gx.DELETE("", keysController.Del)
	// etcd服务列表
	gx.GET("/members", keysController.Members)

}
