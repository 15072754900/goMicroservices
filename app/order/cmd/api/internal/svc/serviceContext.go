package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"look-cp/app/order/cmd/api/internal/config"
	"look-cp/app/order/cmd/rpc/order"
	"look-cp/app/payment/cmd/rpc/payment"
	"look-cp/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	// 之前学习的设计模式里面的一种
	OrderRpc   order.Order
	TravelRpc  travel.Travel
	PaymentRpc payment.Payment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		// 注册rpc服务
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		TravelRpc:  travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
	}
}
