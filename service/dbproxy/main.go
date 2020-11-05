package main

import (
	"github.com/yguilai/go-cloud-storage/common"
	"github.com/yguilai/go-cloud-storage/service/dbproxy/config"
	"github.com/yguilai/go-cloud-storage/service/global"
	"log"
	"time"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"

	dbConn "github.com/yguilai/go-cloud-storage/service/dbproxy/conn"
	dbProxy "github.com/yguilai/go-cloud-storage/service/dbproxy/proto"
	dbRpc "github.com/yguilai/go-cloud-storage/service/dbproxy/rpc"
)

func startRpcService() {
	service := micro.NewService(
		micro.Registry(global.ConsulRegistry),
		micro.Name("go.micro.service.dbproxy"), 	// 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),      	// 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			// 检查是否指定dbhost
			dbhost := c.String("dbhost")
			if len(dbhost) > 0 {
				log.Println("custom db address: " + dbhost)
				config.UpdateDBHost(dbhost)
			}
		}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}
