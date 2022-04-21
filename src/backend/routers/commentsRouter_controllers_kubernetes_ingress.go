package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeIngressController = "kube_web/controllers/kubernetes/ingress:KubeIngressController"
	beego.GlobalControllerRouter[KubeIngressController] = append(
		beego.GlobalControllerRouter[KubeIngressController],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:ingressId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
