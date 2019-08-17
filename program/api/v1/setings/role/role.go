package role

import (
	"net/http"
	"strconv"
	"time"

	"github.com/etcd-manage/etcd-manage-server/program/models"
	"github.com/gin-gonic/gin"
)

/* 角色管理 */

// RoleController 系统管理设置相关
type RoleController struct {
}

// All 获取所有角色
func (s *RoleController) All(c *gin.Context) {
	list, err := new(models.RolesModel).All()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, list)
}

// Add 添加角色
func (s *RoleController) Add(c *gin.Context) {
	req := new(models.RolesModel)
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

// Del 删除角色
func (s *RoleController) Del(c *gin.Context) {
	id := c.Query("id")
	idNum, _ := strconv.Atoi(id)
	if idNum == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "参数错误",
		})
		return
	}
	err := new(models.RolesModel).Del(int32(idNum))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Up 修改角色
func (s *RoleController) Update(c *gin.Context) {
	req := new(models.RolesModel)
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
