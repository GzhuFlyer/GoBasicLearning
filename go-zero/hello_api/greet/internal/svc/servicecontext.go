package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"greet/internal/config"
	"shorturl/rpc/transform"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transform.TransformClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transform.NewTransformClient(zrpc.MustNewClient(c.Transform)),
	}
}
