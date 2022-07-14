package plugins

import (
	"context"
	"fmt"
	"launchbackend/eu.suro/launch/protos/web"
	"launchbackend/internal/handlers"

	"github.com/gin-gonic/gin"
)

type handler struct {
	client web.WebClient
}

func NewPluginHandler(client web.WebClient) handlers.Handler {
	return &handler{
		client: client,
	}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/plugin/:id", h.GetPlugin)
	r.POST("/plugin", h.CreatePlugin)
	r.PUT("/plugin", h.UpdatePlugin)
	r.DELETE("/plugin/:id", h.DeletePlugin)
	r.GET("/plugins", h.GetPlugins)
}

func (h *handler) CreatePlugin(ctx *gin.Context) {
	var dto CreatePluginDTO
	fmt.Print(ctx.Request.Form)
	if er := ctx.ShouldBind(&dto); er != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}

	_, err := h.client.CreatePlugin(context.TODO(), &web.CreatePluginResponse{
		Name:        dto.Name,
		Description: dto.Description,
		Spigotid:    dto.SpigotID,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}

func (h *handler) UpdatePlugin(ctx *gin.Context) {
	var dto UpdatePluginDTO
	fmt.Print(ctx.Request.Form)
	if er := ctx.ShouldBind(&dto); er != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}

	_, err := h.client.UpdatePlugin(context.TODO(), &web.Plugin{
		Id:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Spigotid:    dto.SpigotID,
	})
	if err != nil {
		ctx.JSON(500, gin.H{
			"code":    500,
			"message": "internal server error",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
	})
}

func (h *handler) DeletePlugin(ctx *gin.Context) {
	plugin_id := ctx.Param("id")
	_, err := h.client.DeletePlugin(context.TODO(), &web.DeletePluginResponse{
		Id: plugin_id,
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
}

func (h *handler) GetPlugin(ctx *gin.Context) {
	plugin_id := ctx.Param("id")
	plugin, err := h.client.GetPlugin(context.TODO(), &web.ResponseById{
		Id: plugin_id,
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    plugin,
	})
}

func (h *handler) GetPlugins(ctx *gin.Context) {
	plugins, err := h.client.GetPlugins(context.TODO(), &web.Empty{})
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    plugins,
	})
}
