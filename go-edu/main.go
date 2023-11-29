package main

import (
	"context"
	"flag"
	"go-edu/config"
	"go-edu/logger"
	"go-edu/router"
	"go-edu/tools/jenkins"
	"go-edu/tools/mysql"
	"go-edu/tools/redis"
	"go-edu/util"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	var cfn string
	flag.StringVar(&cfn, "conf", "./conf/config.yaml", "指定配置文件路径")
	flag.Parse()
	// 加载配置文件
	err := config.Init(cfn)
	if err != nil {
		panic(err) // 程序启动时加载配置文件失败直接退出
	}
	// 加载日志
	err = logger.Init(config.Conf.LogConfig, config.Conf.Mode)
	if err != nil {
		panic(err) // 程序启动时初始化日志模块失败直接退出
	}

	//初始化MySQL
	err = mysql.Init(config.Conf.MySQLConfig)
	if err != nil {
		panic(err) // 程序启动时初始化MySQL失败直接退出
	}
	//初始化REDIS
	err = redis.Init(config.Conf.RedisConfig)
	if err != nil {
		panic(err) // 程序启动时初始化REDIS失败直接退出
	}
	//初始化jenkins
	err = jenkins.Init(config.Conf.JenkinsConfig)
	if err != nil {
		panic("程序启动时初始化jenkins失败直接退出") // 程序启动时初始化jenkins失败直接退出
	}

	// 路由初始化
	r := router.Setup()

	// 程序启动
	err = r.Run(config.Conf.IP + ":" + util.InttoStr(config.Conf.Port))
	if err != nil {
		return
	}
	// 等待中断信号以优雅地关闭服务器（设置 15 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Println("Server exiting")
}
