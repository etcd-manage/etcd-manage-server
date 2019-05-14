package common

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/etcd-manage/etcdsdk/model"
	"github.com/gin-gonic/gin"
)

// GetRootDir 获取执行路径
func GetRootDir() string {
	// 文件不存在获取执行路径
	file, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		file = fmt.Sprintf(".%s", string(os.PathSeparator))
	} else {
		file = fmt.Sprintf("%s%s", file, string(os.PathSeparator))
	}
	return file
}

// PathExists 判断文件或目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// GetEtcdClientByGinContext 获取一个etcd客户端 从gin请求上下文
func GetEtcdClientByGinContext(c *gin.Context) (client model.EtcdSdk, err error) {
	clientI, exists := c.Get("CLIENT")
	if exists == false || clientI == nil {
		err = errors.New("Etcd client is empty")
		return
	}
	client = clientI.(model.EtcdSdk)
	return
}
