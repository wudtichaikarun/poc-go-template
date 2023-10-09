package handler

import (
	"net/http"

	"github.com/wudtichaikarun/poc-go-template/api/rest/router"
	"github.com/wudtichaikarun/poc-go-template/service"
)

type CommonHandler struct {
	service.CommonService
}

func (h *CommonHandler) HealthCheck(c router.Context) {
	if ok := h.CommonService.HealthCheck(); !ok {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "service not hearty",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
	return
}
