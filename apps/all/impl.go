package all

import (
	_ "github.com/xiaoweize/ioc-demo/apps/host/http" //注册handler到ioc中
	_ "github.com/xiaoweize/ioc-demo/apps/host/impl" //注册host模块到ioc中
)
