package host

import "context"

//业务接口
type Service interface {
	CreateHost(ctx context.Context, ins *Host) *Host
}
