package main

import (
	"launchbackend/eu.suro/launch/protos/server"
	"launchbackend/eu.suro/launch/protos/web"
	"launchbackend/internal/servers"

	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	serverClient, webClient := startGRPCClient()
	servers := servers.NewServersHandler(serverClient, webClient)
	r := gin.Default()
	servers.Register(r)
}

func startGRPCClient() (serverClient server.ServerClient, webClient web.WebClient) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	serverClient = server.NewServerClient(conn)
	webClient = web.NewWebClient(conn)
	return serverClient, webClient
}
