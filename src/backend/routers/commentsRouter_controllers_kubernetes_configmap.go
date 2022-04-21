package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeConfigMapController = "kube_web/controllers/kubernetes/configmap:KubeConfigMapController"
	beego.GlobalControllerRouter[KubeConfigMapController] = append(
		beego.GlobalControllerRouter[KubeConfigMapController],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:configMapId/tpls/:tplId/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
