package program

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/etcd-manage/etcd-manage-server/program/logger"
	"github.com/etcd-manage/etcd-manage-server/program/models"
	"github.com/etcd-manage/etcdsdk"
	"github.com/etcd-manage/etcdsdk/model"
	"github.com/gin-gonic/autotls"
	gin "github.com/gin-gonic/gin"
)

// http服务

func (p *Program) startAPI() {
	router := gin.Default()

	// 跨域问题
	router.Use(p.middlewareCORS())

	// 设置静态文件目录
	router.GET("/ui/*w", p.handlerStatic)
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/ui")
	})

	// 启动所有版本api
	for key, val := range p.vApis {
		vAPI := router.Group("/" + key)
		vAPI.Use()
		val.Register(vAPI)
	}

	addr := fmt.Sprintf("%s:%d", p.cfg.HTTP.Address, p.cfg.HTTP.Port)
	// 监听
	s := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Println("Start HTTP the service:", addr)
	var err error
	if p.cfg.HTTP.TLSEnable == true {
		if p.cfg.HTTP.TLSConfig == nil || p.cfg.HTTP.TLSConfig.CertFile == "" || p.cfg.HTTP.TLSConfig.KeyFile == "" {
			log.Fatalln("Enable tls must configure certificate file path")
		}
		err = s.ListenAndServeTLS(p.cfg.HTTP.TLSConfig.CertFile, p.cfg.HTTP.TLSConfig.KeyFile)
	} else if p.cfg.HTTP.TLSEncryptEnable == true {
		if len(p.cfg.HTTP.TLSEncryptDomainNames) == 0 {
			log.Fatalln("The domain name list cannot be empty")
		}
		err = autotls.Run(router, p.cfg.HTTP.TLSEncryptDomainNames...)
	} else {
		err = s.ListenAndServe()
	}
	if err != nil {
		log.Fatalln(err)
	}

}

// 跨域中间件
func (p *Program) middlewareCORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, EtcdServerName")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func (p *Program) middlewareEtcdClient() gin.HandlerFunc {
	return func(c *gin.Context) {
		etcdId := c.GetHeader("EtcdID")
		log.Println("当前请求EtcdId", etcdId)
		etcdIdNum, _ := strconv.Atoi(etcdId)
		etcdOne := new(models.EtcdServersModel)
		etcdOne, err := etcdOne.FirstById(int32(etcdIdNum))
		if err != nil {
			logger.Log.Errorw("获取etcd服务信息错误", "EtcdID", etcdId, "err", err)
		}
		// 连接etcd
		cfg := &model.Config{
			Version:   etcdOne.Version,
			Address:   strings.Split(etcdOne.Address, ","),
			TlsEnable: etcdOne.TlsEnable == "true",
			CertFile:  etcdOne.CaFile,
			KeyFile:   etcdOne.KeyFile,
			CaFile:    etcdOne.CaFile,
			Username:  etcdOne.Username,
			Password:  etcdOne.Password,
		}
		client, err := etcdsdk.NewClientByConfig(cfg)
		if err != nil {
			logger.Log.Errorw("连接etcd服务错误", "EtcdID", etcdId, "config", cfg, "err", err)
		}
		c.Set("CLIENT", client)
	}
}
