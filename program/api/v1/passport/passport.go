package passport

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/etcd-manage/etcd-manage-server/program/cache"
	"github.com/etcd-manage/etcd-manage-server/program/common"
	"github.com/etcd-manage/etcd-manage-server/program/logger"
	"github.com/etcd-manage/etcd-manage-server/program/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PassportController 登录退出
type PassportController struct {
}

// LoginReq 登录请求参数
type LoginReq struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (api *PassportController) Login(c *gin.Context) {
	req := new(LoginReq)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if req.Username == "" || req.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	// 查询数据库登录
	req.Password = common.Md5Password(req.Password)
	log.Println(req.Password)
	user := new(models.UsersModel)
	err = user.FirstByUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		logger.Log.Errorw("登录查询用户错误", "err", err)
	}
	user.Password = ""
	user.Token = uuid.New().String()
	// 存储Token
	jsUser, _ := json.Marshal(user)
	key := cache.GetLoginKey(user.Token)
	cache.DefaultMemCache.Set(key, string(jsUser), 7*24*time.Hour)

	c.JSON(http.StatusOK, user)
}
