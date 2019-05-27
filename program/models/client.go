package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/etcd-manage/etcd-manage-server/program/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	client *gorm.DB
)

// InitClient 客户端创建
func InitClient(dbCfg *config.MySQLConfig) (err error) {
	if dbCfg == nil {
		err = errors.New("Config is nil")
		return
	}
	// 拼接连接数据库字符串
	connStr := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",
		dbCfg.User,
		dbCfg.Passwd,
		dbCfg.Address,
		dbCfg.Port,
		dbCfg.DbName)
	// 连接数据库
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return
	}
	if dbCfg.Debug == true {
		db.Debug()
	}
	// 禁用表名多元化
	db.SingularTable(true)
	// 连接池最大连接数
	db.DB().SetMaxIdleConns(dbCfg.MaxIdleConns)
	// 默认打开连接数
	db.DB().SetMaxOpenConns(dbCfg.MaxOpenConns)
	// 开启协程ping MySQL数据库查看连接状态
	go func() {
		for {
			// ping
			err = db.DB().Ping()
			if err != nil {
				log.Println("pingdb error ", err)
			}
			// 间隔5s ping一次
			time.Sleep(30 * time.Second)
		}
	}()

	// 全局变量
	client = db
	return
}
