package handler

import (
	"github.com/wudtichaikarun/poc-go-template/api/rest/app"
	"github.com/wudtichaikarun/poc-go-template/service"
)

type Handler struct {
	CommonHandler
}

func New(s *service.Service) *Handler {
	return &Handler{
		CommonHandler: CommonHandler{s.CommonService},
	}
}

func RegisterRouter(a *app.App, h *Handler) {
	a.GET("/system/health", h.CommonHandler.HealthCheck)
}
