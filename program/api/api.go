package api

import (
	gin "github.com/gin-gonic/gin"
)

// API api 路由注册
type API interface {
	Register(router *gin.RouterGroup)
}
