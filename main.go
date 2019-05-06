package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/etcd-manage/etcd-manage-server/program"
)

func main() {
	// 系统日志显示文件和行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 服务对象
	p, err := program.New()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = p.Run()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 监听退出信号
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, os.Kill) // , syscall.SIGUSR1, syscall.SIGUSR2
	<-c
	log.Println("Exit")
}
