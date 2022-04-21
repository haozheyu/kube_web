package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:deployment/detail/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:deployment/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
		beego.ControllerComments{
			Method:           "UpdateScale",
			Router:           `/:deployment/namespaces/:namespace/clusters/:cluster/updatescale`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
		beego.ControllerComments{
			Method:           "Create",
			Router:           `/:deploymentId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/namespaces/:namespace/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
