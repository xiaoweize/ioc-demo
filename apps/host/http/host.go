package http

import (
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/http/response"
	"github.com/xiaoweize/ioc-demo/apps/host"
)

func (h *Handler) createHost(ctx *gin.Context) {
	var host host.Host
	if err := ctx.Bind(&host); err != nil {
		response.Failed(ctx.Writer, err)
	} else {
		response.Success(ctx.Writer, host)
	}
	//请求中传入context
	h.svc.CreateHost(ctx.Request.Context(), &host)
}
