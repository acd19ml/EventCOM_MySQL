package http

import (
	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc form.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/forms", h.createForm)
}
