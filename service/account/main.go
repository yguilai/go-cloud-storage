package main

import (
	"github.com/micro/go-micro"
	"github.com/yguilai/go-cloud-storage/common"
	"github.com/yguilai/go-cloud-storage/service/account/handler"
	"github.com/yguilai/go-cloud-storage/service/account/proto"
	"github.com/yguilai/go-cloud-storage/service/global"
	"log"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Registry(global.ConsulRegistry),
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
	)
	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
