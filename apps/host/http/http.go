package http

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoweize/ioc-demo/apps"
	"github.com/xiaoweize/ioc-demo/apps/host"
)

//定义Handler来暴露app模块业务
type Handler struct {
	svc host.Service
}

func (h *Handler) Config() {
	//从ioc层获取依赖接口实例对象
	h.svc = apps.GetImpl(h.Name()).(host.Service)
}

// func NewHandler(svc host.Service) *Handler {
// 	return &Handler{
// 		svc: svc,
// 	}
// }

func (h *Handler) Registry(r gin.IRouter) {
	r.POST("/hosts", h.createHost)
}

func (h *Handler) Name() string {
	return host.AppName
}

var (
	handler = &Handler{}
)

func init() {
	//http模块的注册，这个时候还没有初始化
	apps.GinRegistry(handler)
}
