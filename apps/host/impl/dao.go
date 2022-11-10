package impl

import (
	"context"
	"fmt"

	"github.com/xiaoweize/ioc-demo/apps/host"
)

//与数据库的交互操作
func (i *HostServiceImpl) Save(ctx context.Context, ins *host.Host) {
	fmt.Printf("SUCCESS!hostid:%s,hostname:%s\n", ins.Id, ins.Name)
}
