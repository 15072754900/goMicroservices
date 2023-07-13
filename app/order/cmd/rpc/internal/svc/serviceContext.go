package svc

import (
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"look-cp/app/order/cmd/rpc/internal/config"
	"look-cp/app/order/model"
	"look-cp/app/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config
	// 配置rpc服务，在外面的order中配置到服务中，而隔壁的server单纯为了编写服务的函数（还不包括内容）
	AsynqClient *asynq.Client

	TravelRpc travel.Travel

	HomestayOrderModel model.HomestayOrderModel
}

// NewServiceContext 创建一个实例（在外面调用的时候，这也属于是实现一种设计模式）
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AsynqClient:        newAsynqClient(c),
		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
