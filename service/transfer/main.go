package main

import (
	"fmt"
	"github.com/yguilai/go-cloud-storage/service/global"
	"log"
	"time"

	"github.com/yguilai/go-cloud-storage/common"
	"github.com/yguilai/go-cloud-storage/config"
	"github.com/yguilai/go-cloud-storage/mq"
	dbproxy "github.com/yguilai/go-cloud-storage/service/dbproxy/client"
	"github.com/yguilai/go-cloud-storage/service/transfer/process"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
)

func startRPCService() {
	service := micro.NewService(
		micro.Registry(global.ConsulRegistry),
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Flags(common.CustomFlags...),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 检查是否指定mqhost
			mqhost := c.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				mq.UpdateRabbitHost(mqhost)
			}
		}),
	)

	// 初始化 dbproxy client
	dbproxy.Init(service)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTransferService() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")

	// 初始化mq client
	mq.Init()

	mq.StartConsume(
		config.TransOSSQueueName,
		"transfer_oss",
		process.Transfer)
}

func main() {
	// 文件转移服务
	go startTransferService()
	// rpc 服务
	startRPCService()
}
