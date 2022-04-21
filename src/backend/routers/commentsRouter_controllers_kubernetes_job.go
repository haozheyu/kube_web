package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const KubeJobController = "kube_web/controllers/kubernetes/job:KubeJobController"
	beego.GlobalControllerRouter[KubeJobController] = append(
		beego.GlobalControllerRouter[KubeJobController],
		beego.ControllerComments{
			Method:           "ListJobByCronJob",
			Router:           `/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil},
		beego.ControllerComments{
			Method:           "DeleteJob",
			Router:           `/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
