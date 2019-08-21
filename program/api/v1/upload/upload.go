package upload

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadController struct {
}

// UploadOutContent 上传文件，返回内容
func (api *UploadController) UploadOutContent(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
		}
	}()

	fileInfo, err := c.FormFile("file")
	if err != nil {
		return
	}
	if fileInfo.Size > 1*1024*1024 {
		err = errors.New("上传文件大于1Mb")
		return
	}
	f, err := fileInfo.Open()
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}
