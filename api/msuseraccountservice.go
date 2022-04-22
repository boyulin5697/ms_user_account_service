package main

import (
	"flag"
	"fmt"

	"api/internal/config"
	"api/internal/handler"
	"api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/msuseraccountservice-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("成功在 %s:%d启动用户账户微服务...\n", c.Host, c.Port)
	server.Start()
}
