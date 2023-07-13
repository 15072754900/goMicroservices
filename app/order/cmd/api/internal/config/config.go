package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth struct {
		AccessSecret string
	}

	// 注册使用业务层面的rpc服务
	OrderRpcConf   zrpc.RpcClientConf
	PaymentRpcConf zrpc.RpcClientConf
	TravelRpcConf  zrpc.RpcClientConf
}
