package v1

import (
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/keys"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/server"
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
	gx.GET("/val", keysController.Val)
	gx.POST("", keysController.Add)
	gx.PUT("", keysController.Put)
	gx.DELETE("", keysController.Del)

	// 一个etcd服务集群信息
	gx.GET("/members", keysController.Members)

	// etcd服务列表
	serverController := new(server.ServerController)
	gs := router.Group("/server")
	gs.GET("", serverController.List)
	gs.GET("/restore", serverController.Restore)

}
