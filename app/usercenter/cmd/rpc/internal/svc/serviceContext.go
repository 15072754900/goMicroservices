package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"look-cp/app/usercenter/cmd/rpc/internal/config"
	"look-cp/app/usercenter/model"
)

type ServiceContext struct {
	Config config.Config

	RedisClient *redis.Redis

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config: c,

		RedisClient: redis.MustNewRedis(c.Redis.RedisConf, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),

		UserAuthModel: model.NewUserAuthModel(sqlConn, c.Cache),
		UserModel:     model.NewUserModel(sqlConn, c.Cache),
	}
}
