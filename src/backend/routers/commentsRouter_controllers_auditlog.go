package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const AuditLogController = "kube_web/controllers/auditlog:AuditLogController"
	beego.GlobalControllerRouter[AuditLogController] = append(
		beego.GlobalControllerRouter[AuditLogController],
		beego.ControllerComments{
			Method:           "List",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
