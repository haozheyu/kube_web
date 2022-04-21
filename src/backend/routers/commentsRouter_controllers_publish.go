package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const PublishController = "kube_web/controllers/publish:PublishController"
	beego.GlobalControllerRouter[PublishController] = append(
		beego.GlobalControllerRouter[PublishController],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/histories`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "Statistics",
			Router:           `/statistics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
