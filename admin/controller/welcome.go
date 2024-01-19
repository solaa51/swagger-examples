package controller

import (
	"github.com/solaa51/swagger/context"
	"github.com/solaa51/swagger/control"
)

type Welcome struct {
	control.Controller
}

func (w *Welcome) Index(ctx *context.Context) {
	//未知路由 都转发到前端页面
	//f := appPath.AppDir() + appConfig.Info().Static.LocalPath + appConfig.Info().Static.Index
	//http.ServeFile(ctx.ResponseWriter, ctx.Request, f)
}
