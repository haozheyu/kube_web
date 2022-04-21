package cmd

import (
	beego "github.com/beego/beego/v2/server/web"

	"kube_web/initial"
	_ "kube_web/routers"
	"kube_web/workers/webhook"
)

func Run() {

	// MySQL
	initial.InitDb()

	// Swagger API
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// K8S Client
	initial.InitClient()

	// 初始化RsaPrivateKey
	initial.InitRsaKey()

	busEnable := beego.AppConfig.DefaultBool("BusEnable", false)
	if busEnable {
		initial.InitBus()
	}

	webhook.RegisterHookHandler()

	// init kube labels
	initial.InitKubeLabel()

	beego.Run()
}
