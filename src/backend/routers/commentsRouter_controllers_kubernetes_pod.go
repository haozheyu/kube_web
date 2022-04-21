package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubePodController = "kube_web/controllers/kubernetes/pod:KubePodController"
	beego.GlobalControllerRouter[KubePodController] = append(
		beego.GlobalControllerRouter[KubePodController],
		beego.ControllerComments{
			Method:           "Terminal",
			Router:           `/:pod/terminal/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "List",
			Router:           `/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
