package main

import (
	"fmt"
	"github.com/yguilai/go-cloud-storage/service/global"
	"time"

	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	"github.com/yguilai/go-cloud-storage/common"
	dbproxy "github.com/yguilai/go-cloud-storage/service/dbproxy/client"
	cfg "github.com/yguilai/go-cloud-storage/service/download/config"
	dlProto "github.com/yguilai/go-cloud-storage/service/download/proto"
	"github.com/yguilai/go-cloud-storage/service/download/route"
	dlRpc "github.com/yguilai/go-cloud-storage/service/download/rpc"
)

func startRPCService() {
	service := micro.NewService(
		micro.Registry(global.ConsulRegistry),
		micro.Name("go.micro.service.download"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
