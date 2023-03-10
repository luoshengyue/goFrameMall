package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"goFrameMall/internal/controller"
	"goFrameMall/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				//group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().ResponseHandler,
				)
				group.Bind(
					controller.Hello,
					controller.Rotation,     // 轮播图
					controller.Position,     // 手工位
					controller.Admin.Create, //管理员信息
					controller.Admin.Update, //管理员信息
					controller.Admin.Delete, //管理员信息
					controller.Admin.List,   //管理员信息
					controller.Login,
					controller.Data,
					controller.Role,
					controller.Permission,
					controller.Order,
				) //登录
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.ALLMap(g.Map{
						"/backend/admin/info": controller.Admin.Info,
					})
					group.Bind(
						controller.File,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
