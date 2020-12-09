package router

import (
	"entrust/consts"
	"entrust/handler"
	"github.com/gin-gonic/gin"
)

var Engine *gin.Engine

func init() {
	Engine = gin.New()
	// 加载中间件, 如: log, recovery, etc...
	Engine.Use(gin.Recovery())
}

// InitializeRouter 初始化路由
func InitializeRouter() {
	// 路由版本: 初始化为v1
	v1 := Engine.Group(consts.V1)
	// 下单接口参数检查器
	checker := v1.Group(consts.Empty, handler.OrderParamCheck)

	// 下委托单
	checker.POST(consts.Entrust, handler.Entrust)
	// 撤销委托单
	v1.POST(consts.Cancel, handler.Cancel)
}
