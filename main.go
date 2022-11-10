package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoweize/ioc-demo/apps"

	//注册所有模块到ioc中
	_ "github.com/xiaoweize/ioc-demo/apps/all"
)

func main() {
	g := gin.Default()
	//初始化app模块
	apps.ImplInit()
	//初始化http handler模块
	apps.GinInit(g)
	g.Run(":8050")
}
