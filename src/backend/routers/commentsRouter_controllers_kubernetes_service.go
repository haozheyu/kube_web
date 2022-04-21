package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeServiceController = "kube_web/controllers/kubernetes/service:KubeServiceController"
	beego.GlobalControllerRouter[KubeServiceController] = append(
		beego.GlobalControllerRouter[KubeServiceController],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:service/detail/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:serviceId/tpls/:tplId/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
