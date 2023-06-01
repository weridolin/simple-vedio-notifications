package main

import (
	"flag"
	"fmt"

	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/config"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/handler"
	"github.com/weridolin/simple-vedio-notifications/servers/tasks/cmd/rest/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/tasks.yaml", "the config file")

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
