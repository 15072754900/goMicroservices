package main

import (
	"flag"
	"fmt"

	"look-cp/app/usercenter/cmd/api/internal/config"
	"look-cp/app/usercenter/cmd/api/internal/handler"
	"look-cp/app/usercenter/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "C:/Users/86150/Desktop/go-zero-look-cp/app/usercenter/cmd/api/etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
