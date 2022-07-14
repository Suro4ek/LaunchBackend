package version

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

func NewVersionHandler(client web.WebClient) handlers.Handler {
	return &handler{
		client: client,
	}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/version/:id", h.GetVersion)
	r.POST("/version", h.CreateVersion)
	r.PUT("/version", h.UpdateVersion)
	r.DELETE("/version/:id", h.DeleteVersion)
	r.GET("/versions", h.GetVersions)
}

func (h *handler) CreateVersion(ctx *gin.Context) {
	var dto CreateVersionDTO
	// guild_id := ctx.Param("id")
	fmt.Print(ctx.Request.Form)
	if er := ctx.ShouldBind(&dto); er != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}

	_, err := h.client.CreateVersion(context.TODO(), &web.CreateVersionResponse{
		Name:        dto.Name,
		Description: dto.Description,
		Url:         *dto.Url,
		JavaVersion: dto.JavaVersion,
		Version:     dto.Version,
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

func (h *handler) UpdateVersion(ctx *gin.Context) {
	var dto UpdateVersionDTO
	fmt.Print(ctx.Request.Form)
	if er := ctx.ShouldBind(&dto); er != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}

	_, err := h.client.UpdateVersion(context.TODO(), &web.Version{
		Id:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
		Url:         *dto.Url,
		JavaVersion: dto.JavaVersion,
		Version:     dto.Version,
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

func (h *handler) DeleteVersion(ctx *gin.Context) {
	version_id := ctx.Param("id")
	_, err := h.client.DeleteVersion(context.TODO(), &web.DeleteVersionResponse{
		Id: version_id,
	})
	if err != nil {
		ctx.JSON(400, gin.H{
			"code":    400,
			"message": "bad request",
		})
		return
	}
}

func (h *handler) GetVersion(ctx *gin.Context) {
	version_id := ctx.Param("id")
	version, err := h.client.GetVersion(context.TODO(), &web.ResponseById{
		Id: version_id,
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
		"data":    version,
	})
}

func (h *handler) GetVersions(ctx *gin.Context) {
	versions, err := h.client.GetVersions(context.TODO(), &web.Empty{})
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
		"data":    versions,
	})
}
