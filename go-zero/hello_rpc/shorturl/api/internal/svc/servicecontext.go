package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"shorturl/api/internal/config"
	"shorturl/rpc/transformclient"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformclient.Transform
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformclient.NewTransform(zrpc.MustNewClient(c.Transform)),
	}
}
