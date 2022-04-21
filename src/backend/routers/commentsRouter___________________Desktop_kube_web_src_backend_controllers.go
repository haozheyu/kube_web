package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/apikey:ApiKeyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/app:AppController"] = append(beego.GlobalControllerRouter["kube_web/controllers/app:AppController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/appstarred:AppStarredController"] = append(beego.GlobalControllerRouter["kube_web/controllers/appstarred:AppStarredController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/appstarred:AppStarredController"] = append(beego.GlobalControllerRouter["kube_web/controllers/appstarred:AppStarredController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:appId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/auditlog:AuditLogController"] = append(beego.GlobalControllerRouter["kube_web/controllers/auditlog:AuditLogController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"] = append(beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"],
        beego.ControllerComments{
            Method: "CurrentUser",
            Router: "/currentuser",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"] = append(beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/:type/?:name",
            AllowHTTPMethods: []string{"get","post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"] = append(beego.GlobalControllerRouter["kube_web/controllers/auth:AuthController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/bill:BillController"] = append(beego.GlobalControllerRouter["kube_web/controllers/bill:BillController"],
        beego.ControllerComments{
            Method: "ListInvoice",
            Router: "/:appid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/bill:BillController"] = append(beego.GlobalControllerRouter["kube_web/controllers/bill:BillController"],
        beego.ControllerComments{
            Method: "ListCharge",
            Router: "/:appid/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cluster:ClusterController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:BaseConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:BaseConfigController"],
        beego.ControllerComments{
            Method: "ListBase",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"] = append(beego.GlobalControllerRouter["kube_web/controllers/config:ConfigController"],
        beego.ControllerComments{
            Method: "ListSystem",
            Router: "/system",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/configmap:ConfigMapTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/cronjob:CronjobTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:CustomLinkController"],
        beego.ControllerComments{
            Method: "ChangeStatus",
            Router: "/:id/status",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:LinkTypeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/customlink:ShowLinkController"] = append(beego.GlobalControllerRouter["kube_web/controllers/customlink:ShowLinkController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/links",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/daemonset:DaemonSetTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/deployment:DeploymentTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPAController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/hpa:HPATplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/ingress:IngressTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/configmap:KubeConfigMapController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/configmap:KubeConfigMapController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:configMapId/tpls/:tplId/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCRDController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/crd:KubeCustomCRDController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:cronjob/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:cronjob/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:cronjobId/tpls/:tplId/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/cronjob:KubeCronjobController"],
        beego.ControllerComments{
            Method: "Suspend",
            Router: "/:cronjobId/tpls/:tplId/clusters/:cluster/suspend",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:daemonSet/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:daemonSet/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/daemonset:KubeDaemonSetController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:daemonSetId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:deployment/detail/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:deployment/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
        beego.ControllerComments{
            Method: "UpdateScale",
            Router: "/:deployment/namespaces/:namespace/clusters/:cluster/updatescale",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:deploymentId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/deployment:KubeDeploymentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/event:KubeEventController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/event:KubeEventController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/hpa:KubeHPAController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/hpa:KubeHPAController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:hpaId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/ingress:KubeIngressController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/ingress:KubeIngressController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:ingressId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/job:KubeJobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/job:KubeJobController"],
        beego.ControllerComments{
            Method: "DeleteJob",
            Router: "/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/job:KubeJobController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/job:KubeJobController"],
        beego.ControllerComments{
            Method: "ListJobByCronJob",
            Router: "/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/log:KubeLogController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/log:KubeLogController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/:pod/containers/:container/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"],
        beego.ControllerComments{
            Method: "Resources",
            Router: "/:namespaceid([0-9]+)/resources",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/namespace:KubeNamespaceController"],
        beego.ControllerComments{
            Method: "Statistics",
            Router: "/:namespaceid([0-9]+)/statistics",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "AddLabel",
            Router: "/:name/clusters/:cluster/label",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "DeleteLabel",
            Router: "/:name/clusters/:cluster/label",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "GetLabels",
            Router: "/:name/clusters/:cluster/labels",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "AddLabels",
            Router: "/:name/clusters/:cluster/labels",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "DeleteLabels",
            Router: "/:name/clusters/:cluster/labels",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "SetTaint",
            Router: "/:name/clusters/:cluster/taint",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "DeleteTaint",
            Router: "/:name/clusters/:cluster/taint",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/node:KubeNodeController"],
        beego.ControllerComments{
            Method: "NodeStatistics",
            Router: "/statistics",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pod:KubePodController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pod:KubePodController"],
        beego.ControllerComments{
            Method: "Terminal",
            Router: "/:pod/terminal/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pod:KubePodController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pod:KubePodController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/proxy:KubeProxyController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:name/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:KubePersistentVolumeController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:RobinPersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:RobinPersistentVolumeController"],
        beego.ControllerComments{
            Method: "ListRbdImages",
            Router: "/rbd.images/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:RobinPersistentVolumeController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pv:RobinPersistentVolumeController"],
        beego.ControllerComments{
            Method: "CreateRbdImage",
            Router: "/rbd.images/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:KubePersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:KubePersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:pvcId/tpls/:tplId/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "ActiveImage",
            Router: "/:pvc/rbd/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "InActiveImage",
            Router: "/:pvc/rbd/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "DeleteSnapshot",
            Router: "/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "RollbackSnapshot",
            Router: "/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "CreateSnapshot",
            Router: "/:pvc/snapshot/:version/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "ListSnapshot",
            Router: "/:pvc/snapshot/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "DeleteAllSnapshot",
            Router: "/:pvc/snapshot/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "GetPvcStatus",
            Router: "/:pvc/status/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "OfflineImageUser",
            Router: "/:pvc/user/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "LoginInfo",
            Router: "/:pvc/user/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/pvc:RobinPersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Verify",
            Router: "/:pvc/verify/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/secret:KubeSecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/secret:KubeSecretController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:secretId/tpls/:tplId/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/service:KubeServiceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/service:KubeServiceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:service/detail/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/service:KubeServiceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/service:KubeServiceController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:serviceId/tpls/:tplId/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/statefulset:KubeStatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/statefulset:KubeStatefulsetController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:statefulset/namespaces/:namespace/clusters/:cluster",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/kubernetes/statefulset:KubeStatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/kubernetes/statefulset:KubeStatefulsetController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/:statefulsetId([0-9]+)/tpls/:tplId([0-9]+)/clusters/:cluster",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Statistics",
            Router: "/:id([0-9]+)/statistics",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "History",
            Router: "/:id/history",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "InitDefault",
            Router: "/init",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "Migrate",
            Router: "/migration",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"] = append(beego.GlobalControllerRouter["kube_web/controllers/namespace:NamespaceController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"] = append(beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"] = append(beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"] = append(beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"],
        beego.ControllerComments{
            Method: "Publish",
            Router: "/",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"] = append(beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"],
        beego.ControllerComments{
            Method: "Subscribe",
            Router: "/subscribe",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"] = append(beego.GlobalControllerRouter["kube_web/controllers/notification:NotificationController"],
        beego.ControllerComments{
            Method: "Read",
            Router: "/subscribe",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "ListAppDeploys",
            Router: "/get_app_deploys",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetDeploymentDetail",
            Router: "/get_deployment_detail",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetDeploymentStatus",
            Router: "/get_deployment_status",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetLatestDeploymentTpl",
            Router: "/get_latest_deployment_tpl",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetPodInfo",
            Router: "/get_pod_info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetPodInfoFromIP",
            Router: "/get_pod_info_from_ip",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetPodList",
            Router: "/get_pod_list",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "GetResourceInfo",
            Router: "/get_resource_info",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "ListNamespaceApps",
            Router: "/list_namespace_apps",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "ListNamespaceUsers",
            Router: "/list_namespace_users",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "RestartDeployment",
            Router: "/restart_deployment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "ScaleDeployment",
            Router: "/scale_deployment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"] = append(beego.GlobalControllerRouter["kube_web/controllers/openapi:OpenAPIController"],
        beego.ControllerComments{
            Method: "UpgradeDeployment",
            Router: "/upgrade_deployment",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:AppUserController"],
        beego.ControllerComments{
            Method: "GetPermissionByApp",
            Router: "/permissions/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:GroupController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:NamespaceUserController"],
        beego.ControllerComments{
            Method: "GetPermissionByNS",
            Router: "/permissions/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:PermissionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "UpdateAdmin",
            Router: "/:id/admin",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "ResetPassword",
            Router: "/:id/resetpassword",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"] = append(beego.GlobalControllerRouter["kube_web/controllers/permission:UserController"],
        beego.ControllerComments{
            Method: "UserStatistics",
            Router: "/statistics",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/publish:PublishController"] = append(beego.GlobalControllerRouter["kube_web/controllers/publish:PublishController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/histories",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/publish:PublishController"] = append(beego.GlobalControllerRouter["kube_web/controllers/publish:PublishController"],
        beego.ControllerComments{
            Method: "Statistics",
            Router: "/statistics",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/publishstatus:PublishStatusController"] = append(beego.GlobalControllerRouter["kube_web/controllers/publishstatus:PublishStatusController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/publishstatus:PublishStatusController"] = append(beego.GlobalControllerRouter["kube_web/controllers/publishstatus:PublishStatusController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/pvc:PersistentVolumeClaimTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/secret:SecretTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "GetNames",
            Router: "/names",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetController"],
        beego.ControllerComments{
            Method: "UpdateOrders",
            Router: "/updateorders",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"] = append(beego.GlobalControllerRouter["kube_web/controllers/statefulset:StatefulsetTplController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "Create",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "Update",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:id([0-9]+)",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"] = append(beego.GlobalControllerRouter["kube_web/controllers/webhook:WebHookController"],
        beego.ControllerComments{
            Method: "GetHookEvents",
            Router: "/events",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
