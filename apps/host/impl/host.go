package impl

import (
	"context"

	"github.com/xiaoweize/ioc-demo/apps/host"
)

//业务接口实现
func (h *HostServiceImpl) CreateHost(ctx context.Context, ins *host.Host) *host.Host {
	h.Save(ctx, ins)
	return nil
}
