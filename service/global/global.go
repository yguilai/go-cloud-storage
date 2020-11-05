package global

import (
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
)

var (
	ConsulRegistry = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.88.131:8500",
		}
	})
)
