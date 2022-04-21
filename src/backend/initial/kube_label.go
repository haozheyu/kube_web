package initial

import (
	beego "github.com/beego/beego/v2/server/web"

	"kube_web/util"
)

func InitKubeLabel() {
	util.AppLabelKey = beego.AppConfig.DefaultString("AppLabelKey", "kube-app")
	util.NamespaceLabelKey = beego.AppConfig.DefaultString("NamespaceLabelKey", "kube-app-ns")
	util.PodAnnotationControllerKindLabelKey = beego.AppConfig.DefaultString("PodAnnotationControllerKindLabelKey", "kube-app/controller-kind")
}
