package main

import (
	"flag"
	"fmt"
	"look-cp/common/interceptor/rpcserver"

	"look-cp/app/order/cmd/rpc/internal/config"
	"look-cp/app/order/cmd/rpc/internal/server"
	"look-cp/app/order/cmd/rpc/internal/svc"
	"look-cp/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "C:/Users/86150/Desktop/go-zero-look-cp/app/order/cmd/rpc/etc/order.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewOrderServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterOrderServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// rpc log
	s.AddUnaryInterceptors(rpcserver.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
