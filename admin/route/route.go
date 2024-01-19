package route

import (
	"github.com/solaa51/swagger-examples/admin/controller"
	"github.com/solaa51/swagger/control"
	"github.com/solaa51/swagger/handle"
	"github.com/solaa51/swagger/router"
)

func init() {
	//配置url匹配规则
	//配置中间件规则

	// 过滤其他路由
	router.AddCompile(`\w.*`, "welcome/index")
}

func init() {
	//配置控制器
	handle.AddHandleStruct(func() control.Control { return &controller.Welcome{} }, "")
}
