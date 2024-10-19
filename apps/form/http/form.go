package http

import (
	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/mcube/http/response"
	"github.com/gin-gonic/gin"
)

// 用于暴露Form service接口
func (h *Handler) createForm(c *gin.Context) {
	ins := form.NewForm()
	if err := c.Bind(ins); err != nil {
		// 参数绑定失败
		response.Failed(c.Writer, err)
		return
	}

	// 调用服务
	ins, err := h.svc.CreateForm(c.Request.Context(), ins)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, ins)
}
