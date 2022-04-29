package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeEventController = "kube_web/controllers/kubernetes/event:KubeEventController"
	beego.GlobalControllerRouter[KubeEventController] = append(
		beego.GlobalControllerRouter[KubeEventController],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/namespaces/clusters/event`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
