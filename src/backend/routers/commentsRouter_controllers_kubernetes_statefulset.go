package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeStatefulsetController = "kube_web/controllers/kubernetes/statefulset:KubeStatefulsetController"
	beego.GlobalControllerRouter[KubeStatefulsetController] = append(
		beego.GlobalControllerRouter[KubeStatefulsetController],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:statefulset/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:statefulsetId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
