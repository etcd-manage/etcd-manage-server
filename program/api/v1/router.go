package v1

import (
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/keys"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/passport"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/server"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/setings/role"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/setings/user"
	"github.com/etcd-manage/etcd-manage-server/program/api/v1/upload"
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
	gs.POST("", serverController.Add)
	gs.PUT("", serverController.Update)
	gs.DELETE("", serverController.Del)
	gs.GET("/restore", serverController.Restore)
	gs.POST("/roles", serverController.SetRoles)
	gs.GET("/roles", serverController.GetRoles)

	// 认证中心
	passportController := new(passport.PassportController)
	gp := router.Group("/passport")
	gp.POST("/login", passportController.Login)

	// 角色
	roleController := new(role.RoleController)
	rs := router.Group("/role")
	rs.GET("", roleController.All)
	rs.POST("", roleController.Add)
	rs.PUT("", roleController.Update)
	rs.DELETE("", roleController.Del)

	// 用户
	userController := new(user.UserController)
	us := router.Group("/user")
	us.GET("", userController.List)
	us.POST("", userController.Add)
	us.PUT("", userController.Update)
	us.DELETE("", userController.Del)

	// 文件上传
	uploadController := new(upload.UploadController)
	uus := router.Group("/upload")
	uus.POST("/content", uploadController.UploadOutContent)

}
