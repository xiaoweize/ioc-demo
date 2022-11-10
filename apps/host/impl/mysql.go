package impl

import (
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"github.com/xiaoweize/ioc-demo/apps"
	"github.com/xiaoweize/ioc-demo/apps/host"
)

var (
	impl = new(HostServiceImpl)
)

//业务接口具体实现
type HostServiceImpl struct {
	l logger.Logger
}

//初始化方法
func (i *HostServiceImpl) Config() {
	i.l = zap.L().Named("host")
}

func (i *HostServiceImpl) Name() string {
	return host.AppName
}

func init() {
	//app模块的注册，这个时候还没有初始化
	apps.ImplRegistry(impl)
}
