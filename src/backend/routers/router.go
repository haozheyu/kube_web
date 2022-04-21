// @APIVersion 1.0.0
// @Title Wayne API
// @Description List接口查询技巧：<br> 1.按照字段排序：增加sortby查询参数。例如：sortby=-id ，负号表示倒叙排序，不加负号表示默认排序 <br> 2.按照字段查询：增加filter查询参数。例如：filter=name__exact=test 表示name名称等于test， filter=name__contains=test表示名称包含test
package routers

import (
	"net/http"
	"os"
	"path"
	"runtime"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"kube_web/controllers/apikey"
	"kube_web/controllers/app"
	"kube_web/controllers/appstarred"
	"kube_web/controllers/auditlog"
	"kube_web/controllers/auth"
	"kube_web/controllers/bill"
	"kube_web/controllers/cluster"
	"kube_web/controllers/config"
	"kube_web/controllers/configmap"
	"kube_web/controllers/cronjob"
	"kube_web/controllers/customlink"
	"kube_web/controllers/daemonset"
	"kube_web/controllers/deployment"
	"kube_web/controllers/hpa"
	"kube_web/controllers/ingress"
	kconfigmap "kube_web/controllers/kubernetes/configmap"
	kcrd "kube_web/controllers/kubernetes/crd"
	kcronjob "kube_web/controllers/kubernetes/cronjob"
	kdaemonset "kube_web/controllers/kubernetes/daemonset"
	kdeployment "kube_web/controllers/kubernetes/deployment"
	kevent "kube_web/controllers/kubernetes/event"
	khpa "kube_web/controllers/kubernetes/hpa"
	kingress "kube_web/controllers/kubernetes/ingress"
	kjob "kube_web/controllers/kubernetes/job"
	klog "kube_web/controllers/kubernetes/log"
	knamespace "kube_web/controllers/kubernetes/namespace"
	knode "kube_web/controllers/kubernetes/node"
	kpod "kube_web/controllers/kubernetes/pod"
	"kube_web/controllers/kubernetes/proxy"
	kpv "kube_web/controllers/kubernetes/pv"
	kpvc "kube_web/controllers/kubernetes/pvc"
	ksecret "kube_web/controllers/kubernetes/secret"
	kservice "kube_web/controllers/kubernetes/service"
	kstatefulset "kube_web/controllers/kubernetes/statefulset"
	"kube_web/controllers/namespace"
	"kube_web/controllers/notification"
	"kube_web/controllers/openapi"
	"kube_web/controllers/permission"
	"kube_web/controllers/publish"
	"kube_web/controllers/publishstatus"
	"kube_web/controllers/pvc"
	"kube_web/controllers/secret"
	"kube_web/controllers/statefulset"
	"kube_web/controllers/webhook"
	"kube_web/health"
	"kube_web/util/hack"
)

func init() {
	// Beego注解路由代码生成规则和程序运行路径相关，需要改写一下避免产生不一致的文件名
	if beego.BConfig.RunMode == "dev" && path.Base(beego.AppPath) == "_build" {
		beego.AppPath = path.Join(path.Dir(beego.AppPath), "src/backend")
	}

	// linux 的go run 执行路径 是/tmp/go-buildxxxx,会使注解路由生成文件名很奇怪
	if beego.BConfig.RunMode == "dev" && runtime.GOOS == "linux" {
		beego.AppPath, _ = os.Getwd()
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		// 其中Options跨域复杂请求预检
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "token", "Access-Control-Allow-Origin", "X-Requested-With", "Accept", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
		// 指定可访问域名AllowOrigins
		AllowOrigins: []string{"*"},
	}))

	beego.Handler("/ws/pods/exec", kpod.CreateAttachHandler("/ws/pods/exec"), true)

	beego.Get("/healthz", func(ctx *context.Context) {
		dc := health.DatabaseCheck{}
		err := dc.Check()
		if err != nil {
			ctx.Output.SetStatus(http.StatusInternalServerError)
			ctx.Output.Body(hack.Slice(err.Error()))
			return
		}
		ctx.Output.SetStatus(http.StatusOK)
		ctx.Output.Body(hack.Slice("ok"))
	})

	beego.Include(&auth.AuthController{})

	nsWithApp := beego.NewNamespace("/api/v1",
		// 路由中携带appid
		beego.NSNamespace("/apps/:appid([0-9]+)/users",
			beego.NSInclude(
				&permission.AppUserController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/configmaps",
			beego.NSInclude(
				&configmap.ConfigMapController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/configmaps/tpls",
			beego.NSInclude(
				&configmap.ConfigMapTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/cronjobs",
			beego.NSInclude(
				&cronjob.CronjobController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/cronjobs/tpls",
			beego.NSInclude(
				&cronjob.CronjobTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/deployments",
			beego.NSInclude(
				&deployment.DeploymentController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/deployments/tpls",
			beego.NSInclude(
				&deployment.DeploymentTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/statefulsets",
			beego.NSInclude(
				&statefulset.StatefulsetController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/statefulsets/tpls",
			beego.NSInclude(
				&statefulset.StatefulsetTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/daemonsets",
			beego.NSInclude(
				&daemonset.DaemonSetController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/daemonsets/tpls",
			beego.NSInclude(
				&daemonset.DaemonSetTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/persistentvolumeclaims",
			beego.NSInclude(
				&pvc.PersistentVolumeClaimController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/persistentvolumeclaims/tpls",
			beego.NSInclude(
				&pvc.PersistentVolumeClaimTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/secrets",
			beego.NSInclude(
				&secret.SecretController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/secrets/tpls",
			beego.NSInclude(
				&secret.SecretTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/webhooks",
			beego.NSInclude(
				&webhook.WebHookController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/apikeys",
			beego.NSInclude(
				&apikey.ApiKeyController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/ingresses",
			beego.NSInclude(
				&ingress.IngressController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/ingresses/tpls",
			beego.NSInclude(
				&ingress.IngressTplController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/hpas",
			beego.NSInclude(
				&hpa.HPAController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/hpas/tpls",
			beego.NSInclude(
				&hpa.HPATplController{},
			),
		),
	)

	nsWithKubernetes := beego.NewNamespace("/api/v1",
		beego.NSRouter("/kubernetes/pods/statistics", &kpod.KubePodController{}, "get:PodStatistics"),

		beego.NSNamespace("/kubernetes/persistentvolumes",
			beego.NSInclude(
				&kpv.KubePersistentVolumeController{},
			),
		),
		beego.NSNamespace("/kubernetes/persistentvolumes/robin",
			beego.NSInclude(
				&kpv.RobinPersistentVolumeController{},
			),
		),
		beego.NSNamespace("/kubernetes/namespaces",
			beego.NSInclude(
				&knamespace.KubeNamespaceController{},
			),
		),
		beego.NSNamespace("/kubernetes/nodes",
			beego.NSInclude(
				&knode.KubeNodeController{},
			),
		),
	)

	nsWithKubernetesApp := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/cronjobs",
			beego.NSInclude(
				&kcronjob.KubeCronjobController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/deployments",
			beego.NSInclude(
				&kdeployment.KubeDeploymentController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/statefulsets",
			beego.NSInclude(
				&kstatefulset.KubeStatefulsetController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/daemonsets",
			beego.NSInclude(
				&kdaemonset.KubeDaemonSetController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/configmaps",
			beego.NSInclude(
				&kconfigmap.KubeConfigMapController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/services",
			beego.NSInclude(
				&kservice.KubeServiceController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/ingresses",
			beego.NSInclude(
				&kingress.KubeIngressController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/hpas",
			beego.NSInclude(
				&khpa.KubeHPAController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/secrets",
			beego.NSInclude(
				&ksecret.KubeSecretController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/persistentvolumeclaims",
			beego.NSInclude(
				&kpvc.KubePersistentVolumeClaimController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/persistentvolumeclaims/robin",
			beego.NSInclude(
				&kpvc.RobinPersistentVolumeClaimController{},
			),
		),

		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/jobs",
			beego.NSInclude(
				&kjob.KubeJobController{},
			),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/pods",
			beego.NSInclude(
				&kpod.KubePodController{}),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/events",
			beego.NSInclude(
				&kevent.KubeEventController{}),
		),
		beego.NSNamespace("/kubernetes/apps/:appid([0-9]+)/podlogs",
			beego.NSInclude(
				&klog.KubeLogController{}),
		),
	)

	nsWithNamespace := beego.NewNamespace("/api/v1",
		// 路由中携带namespaceid
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/apps",
			beego.NSInclude(
				&app.AppController{},
			),
		),
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/webhooks",
			beego.NSInclude(
				&webhook.WebHookController{},
			),
		),
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/apikeys",
			beego.NSInclude(
				&apikey.ApiKeyController{},
			),
		),
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/users",
			beego.NSInclude(
				&permission.NamespaceUserController{},
			),
		),
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/bills",
			beego.NSInclude(
				&bill.BillController{},
			),
		),
		beego.NSNamespace("/namespaces/:namespaceid([0-9]+)/customlink",
			beego.NSInclude(
				&customlink.ShowLinkController{},
			),
		),
	)

	nsWithoutApp := beego.NewNamespace("/api/v1",
		// 路由中不携带任何id
		beego.NSNamespace("/configs",
			beego.NSInclude(
				&config.ConfigController{},
			),
		),
		beego.NSNamespace("/configs/base",
			beego.NSInclude(
				&config.BaseConfigController{},
			),
		),
		beego.NSNamespace("/linktypes",
			beego.NSInclude(
				&customlink.LinkTypeController{},
			),
		),
		beego.NSNamespace("/customlinks",
			beego.NSInclude(
				&customlink.CustomLinkController{},
			),
		),
		beego.NSRouter("/apps/statistics", &app.AppController{}, "get:AppStatistics"),
		beego.NSNamespace("/clusters",
			beego.NSInclude(
				&cluster.ClusterController{},
			),
		),
		beego.NSNamespace("/auditlogs",
			beego.NSInclude(
				&auditlog.AuditLogController{},
			),
		),
		beego.NSNamespace("/notifications",
			beego.NSInclude(
				&notification.NotificationController{},
			),
		),
		beego.NSNamespace("/namespaces",
			beego.NSInclude(
				&namespace.NamespaceController{},
			),
		),
		beego.NSNamespace("/apps/stars",
			beego.NSInclude(
				&appstarred.AppStarredController{},
			),
		),
		beego.NSNamespace("/publish",
			beego.NSInclude(
				&publish.PublishController{},
			),
		),
		beego.NSNamespace("/publishstatus",
			beego.NSInclude(
				&publishstatus.PublishStatusController{},
			),
		),
		beego.NSNamespace("/users",
			beego.NSInclude(
				&permission.UserController{}),
		),
		beego.NSNamespace("/groups",
			beego.NSInclude(
				&permission.GroupController{},
			),
		),
		beego.NSNamespace("/permissions",
			beego.NSInclude(
				&permission.PermissionController{},
			),
		),
	)

	nsWithOpenAPI := beego.NewNamespace("/openapi/v1",
		beego.NSNamespace("/gateway/action",
			beego.NSInclude(
				&openapi.OpenAPIController{}),
		),
	)

	// For Kubernetes resource router
	// appid used to check permission
	nsWithKubernetesProxy := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/namespaces/:namespace/:kind",
			beego.NSInclude(
				&proxy.KubeProxyController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/customresourcedefinitions",
			beego.NSInclude(
				&kcrd.KubeCRDController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/apis/:group/:version/namespaces/:namespace/:kind",
			beego.NSInclude(
				&kcrd.KubeCustomCRDController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/apis/:group/:version/:kind",
			beego.NSInclude(
				&kcrd.KubeCustomCRDController{},
			),
		),
		beego.NSNamespace("/apps/:appid([0-9]+)/_proxy/clusters/:cluster/:kind",
			beego.NSInclude(
				&proxy.KubeProxyController{},
			),
		),
	)

	beego.AddNamespace(nsWithKubernetes)

	beego.AddNamespace(nsWithKubernetesApp)

	beego.AddNamespace(nsWithApp)

	beego.AddNamespace(nsWithoutApp)

	beego.AddNamespace(nsWithNamespace)

	beego.AddNamespace(nsWithOpenAPI)

	beego.AddNamespace(nsWithKubernetesProxy)
}
