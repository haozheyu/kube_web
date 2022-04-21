package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeSecretController = "kube_web/controllers/kubernetes/secret:KubeSecretController"
	beego.GlobalControllerRouter[KubeSecretController] = append(
		beego.GlobalControllerRouter[KubeSecretController],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:secretId/tpls/:tplId/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
