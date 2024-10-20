package http

import (
	"github.com/acd19ml/EventCOM_MySQL/apps"
	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/gin-gonic/gin"
)

func NewFormHTTPHandler() *Handler {
	return &Handler{}
}

var handler = &Handler{}

type Handler struct {
	svc form.Service
}

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/forms", h.createForm)
}

func (h *Handler) Config() {
	if apps.FormService == nil {
		panic("FormService required")
	}

	// 从ioc获取FormService的实例，代替原来从构造函数传入
	h.svc = apps.GetImpl(form.AppName).(form.Service)
}

func (h *Handler) Name() string {
	return form.AppName
}

func init() {
	apps.RegistryGin(handler)
}
