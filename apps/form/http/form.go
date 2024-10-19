package http

import (
	"log"

	"github.com/acd19ml/EventCOM_MySQL/apps/form"
	"github.com/acd19ml/EventCOM_MySQL/mcube/http/response"
	"github.com/gin-gonic/gin"
)

// 用于暴露Form service接口
func (h *Handler) createForm(c *gin.Context) {
	ins := form.NewForm()

	log.Println("Received POST request to create form")
	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
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

	// 成功, 把对象实例返回给HTTP API调用方
	response.Success(c.Writer, ins)
}
