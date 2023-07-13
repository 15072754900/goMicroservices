package svc

import (
	"github.com/hibiken/asynq"
	"look-cp/app/order/cmd/rpc/internal/config"
)

func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Host, Password: c.Redis.Pass})
}
