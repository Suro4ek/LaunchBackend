package servers

import (
	"context"
	"fmt"
	"launchbackend/eu.suro/launch/protos/server"
	"launchbackend/eu.suro/launch/protos/web"
	"launchbackend/internal/handlers"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	serverClient server.ServerClient
	webClient    web.WebClient
}

func NewServersHandler(serverClient server.ServerClient, webClient web.WebClient) handlers.Handler {
	return &handler{
		serverClient: serverClient,
		webClient:    webClient,
	}
}

func (h *handler) Register(router *gin.Engine) {
	router.GET("/servers", h.getServers)
	router.DELETE("/server/:port", h.stopServer)
	router.DELETE("/servers", h.stopAllServers)
}

func (h *handler) getServers(ctx *gin.Context) {
	servers, err := h.serverClient.ListServers(context.TODO(), &server.Empty{})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, servers)
}

func (h *handler) stopServer(ctx *gin.Context) {
	port := ctx.Param("port")
	//string to int32
	portInt, err := strconv.ParseInt(port, 10, 32)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	_, err = h.serverClient.DeleteServer(context.TODO(), &server.DeleteServerRequest{Port: int32(portInt)})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, "Server stopped")
}

func (h *handler) stopAllServers(ctx *gin.Context) {
	_, err := h.webClient.DeleteAllServers(context.TODO(), &web.Empty{})
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, "All servers stopped")
}
