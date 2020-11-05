package main

import (
	"fmt"
	"github.com/yguilai/go-cloud-storage/config"
	"github.com/yguilai/go-cloud-storage/mq"
	"github.com/yguilai/go-cloud-storage/service/global"
	"log"
	"os"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/yguilai/go-cloud-storage/common"
	dbproxy "github.com/yguilai/go-cloud-storage/service/dbproxy/client"
	cfg "github.com/yguilai/go-cloud-storage/service/upload/config"
	upProto "github.com/yguilai/go-cloud-storage/service/upload/proto"
	"github.com/yguilai/go-cloud-storage/service/upload/route"
	upRpc "github.com/yguilai/go-cloud-storage/service/upload/rpc"
)

func startRPCService() {
	service := micro.NewService(
		micro.Registry(global.ConsulRegistry),
		micro.Name("go.micro.service.upload"), // 服务名称
		micro.RegisterTTL(time.Second*10),     // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5), // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
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

	// 初始化dbproxy client
	dbproxy.Init(service)
	// 初始化mq client
	mq.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
}

func main() {
	os.MkdirAll(config.TempLocalRootDir, 0777)
	os.MkdirAll(config.TempPartRootDir, 0777)

	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
