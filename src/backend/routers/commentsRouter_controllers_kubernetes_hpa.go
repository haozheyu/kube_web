package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeHPAController = "kube_web/controllers/kubernetes/hpa:KubeHPAController"
	beego.GlobalControllerRouter[KubeHPAController] = append(
		beego.GlobalControllerRouter[KubeHPAController],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:hpaId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
