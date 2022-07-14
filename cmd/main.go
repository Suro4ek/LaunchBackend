package main

import (
	"launchbackend/eu.suro/launch/protos/server"
	"launchbackend/eu.suro/launch/protos/web"
	"launchbackend/internal/config"
	"launchbackend/internal/plugins"
	"launchbackend/internal/servers"
	"launchbackend/internal/version"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetConfig()
	serverClient, webClient := startGRPCClient(cfg)
	servers := servers.NewServersHandler(serverClient, webClient)
	version := version.NewVersionHandler(webClient)
	plugin := plugins.NewPluginHandler(webClient)
	r := gin.Default()
	servers.Register(r)
	version.Register(r)
	plugin.Register(r)
	r.Run("0.0.0.0:" + cfg.Port)
}

func startGRPCClient(cfg *config.Config) (serverClient server.ServerClient, webClient web.WebClient) {
	conn, err := grpc.Dial(cfg.GRPCHost+":"+cfg.GRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	serverClient = server.NewServerClient(conn)
	webClient = web.NewWebClient(conn)
	return serverClient, webClient
}
