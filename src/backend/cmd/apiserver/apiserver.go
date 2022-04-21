package apiserver

//noinspection ALL
import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/spf13/cobra"

	"kube_web/initial"
	_ "kube_web/routers"
	"kube_web/workers/webhook"
)

var (
	APIServerCmd = &cobra.Command{
		Use:    "apiserver",
		PreRun: preRun,
		Run:    run,
	}
)

func preRun(cmd *cobra.Command, args []string) {
}

func run(cmd *cobra.Command, args []string) {
	// MySQL
	initial.InitDb()

	// Swagger API
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// K8S Client
	initial.InitClient()

	// 初始化RabbitMQ
	busEnable := beego.AppConfig.DefaultBool("BusEnable", false)
	if busEnable {
		initial.InitBus()
	}
	// register webhook event handler
	hookenable := beego.AppConfig.DefaultBool("EnableWebhook", false)
	if hookenable {
		webhook.RegisterHookHandler()
	}

	// 初始化RsaPrivateKey
	initial.InitRsaKey()

	// init kube labels
	initial.InitKubeLabel()

	beego.Run()
}
