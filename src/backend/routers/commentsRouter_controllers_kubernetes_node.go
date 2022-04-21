package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/:name/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "Update",
			Router:           `/:name/clusters/:cluster`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:name/clusters/:cluster`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "AddLabel",
			Router:           `/:name/clusters/:cluster/label`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "DeleteLabel",
			Router:           `/:name/clusters/:cluster/label`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "GetLabels",
			Router:           `/:name/clusters/:cluster/labels`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "AddLabels",
			Router:           `/:name/clusters/:cluster/labels`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "DeleteLabels",
			Router:           `/:name/clusters/:cluster/labels`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "SetTaint",
			Router:           `/:name/clusters/:cluster/taint`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "DeleteTaint",
			Router:           `/:name/clusters/:cluster/taint`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/clusters/:cluster`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
		beego.ControllerComments{
			Method:           "NodeStatistics",
			Router:           `/statistics`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
