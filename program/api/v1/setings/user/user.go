package user

import (
	"net/http"
	"strconv"
	"time"

	"github.com/etcd-manage/etcd-manage-server/program/common"
	"github.com/etcd-manage/etcd-manage-server/program/logger"
	"github.com/etcd-manage/etcd-manage-server/program/models"
	"github.com/gin-gonic/gin"
)

/* 用户管理 */

// UserController 用户管理
type UserController struct {
}

// List 分页获取
func (s *UserController) List(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()
	// 读取条件
	name := c.Query("name")
	userId := common.GetHttpToInt(c, "user_id")
	roleId := common.GetHttpToInt(c, "role_id")
	page := common.GetHttpToInt(c, "page")
	pageSize := common.GetHttpToInt(c, "page_size")
	if page == 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	// 查询列表
	user := new(models.UsersModel)
	list, err := user.List(int32(userId), int32(roleId), name, offset, pageSize)
	if err != nil {
		logger.Log.Errorw("查询用户列表错误", "err", err)
		return
	}
	_c, err := user.ListCount(int32(userId), int32(roleId), name, offset, pageSize)
	if err != nil {
		logger.Log.Errorw("查询用户总数错误", "err", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"list":  list,
		"total": _c,
	})
}

// Add 添加
func (s *UserController) Add(c *gin.Context) {
	req := new(models.UsersModel)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	now := models.JSONTime(time.Now())
	req.CreatedAt = now
	req.UpdatedAt = now
	req.Id = 0
	err = req.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}

// Del 删除
func (s *UserController) Del(c *gin.Context) {
	id := c.Query("id")
	idNum, _ := strconv.Atoi(id)
	if idNum == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	err := new(models.UsersModel).Del(int32(idNum))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Up 修改
func (s *UserController) Update(c *gin.Context) {
	req := new(models.UsersModel)
	err := c.Bind(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if req.Id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	now := models.JSONTime(time.Now())
	req.UpdatedAt = now
	err = req.Save()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, nil)
}
