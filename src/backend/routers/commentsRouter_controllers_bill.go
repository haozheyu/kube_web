package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {
	const BillController = "kube_web/controllers/bill:BillController"
	beego.GlobalControllerRouter[BillController] = append(
		beego.GlobalControllerRouter[BillController],
		beego.ControllerComments{
			Method:           "ListInvoice",
			Router:           `/:appid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		},
		beego.ControllerComments{
			Method:           "ListCharge",
			Router:           `/:appid/:name`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil,
		})
}
