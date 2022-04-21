package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeLogController = "kube_web/controllers/kubernetes/log:KubeLogController"
	beego.GlobalControllerRouter[KubeLogController] = append(
		beego.GlobalControllerRouter[KubeLogController],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/:pod/containers/:container/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
