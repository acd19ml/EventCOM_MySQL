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

func (h *Handler) queryForm(c *gin.Context) {

	req := form.NewQueryFormFromHTTP(c.Request)

	set, err := h.svc.QueryForm(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	// 成功, 把对象实例返回给HTTP API调用方
	response.Success(c.Writer, set)
}

func (h *Handler) describeForm(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := form.NewDescribeFormRequestWithId(c.Param("id"))

	// 进行接口调用, 返回 肯定有成功或者失败
	set, err := h.svc.DescribeForm(c.Request.Context(), req)
	if err != nil {
		response.Failed(c.Writer, err)
		return
	}

	response.Success(c.Writer, set)
}
