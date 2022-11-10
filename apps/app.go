package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	implApps = map[string]ImplService{}
	//注册gin编写的handler
	GinSvcs = map[string]GinService{}
)

type GinService interface {
	Config()              //初始化
	Registry(gin.IRouter) //handler路由注册
	Name() string
}

func GinRegistry(svc GinService) {
	if v, ok := GinSvcs[svc.Name()]; ok {
		panic(fmt.Sprintf("Ginservice %s has registried!", v))
	}
	GinSvcs[svc.Name()] = svc
}

func GinInit(r gin.IRouter) {
	//初始化
	for _, v := range GinSvcs {
		v.Config()
	}
	//注册路由
	for _, v := range GinSvcs {
		v.Registry(r)
	}
}

type ImplService interface {
	Config() //app模块的初始化
	Name() string
}

func ImplRegistry(svc ImplService) {
	if _, ok := implApps[svc.Name()]; ok {
		fmt.Println(implApps)
		panic(fmt.Sprintf("ImplService [%s] has registried!", svc.Name()))
	}
	implApps[svc.Name()] = svc
}

func ImplInit() {
	for _, v := range implApps {
		v.Config()
	}
}

func GetImpl(name string) any {
	for k, v := range implApps {
		if k == name {
			return v
		}
	}
	return nil
}
