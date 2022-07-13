package plugins

import (
	"github.com/gin-gonic/gin"
	"launchbackend/eu.suro/launch/protos/web"
	"launchbackend/internal/handlers"
)

type handler struct {
	webClient web.WebClient
}

func NewPluginsHandler(web web.WebClient) handlers.Handler {
	return &handler{
		web,
	}
}

func (h *handler) Register(r *gin.Engine) {

}

func (h *handler) CreatePlugin(ctx *gin.Context) {
	h.webClient.
}
