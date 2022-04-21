package initial

import (
	"time"

	"k8s.io/apimachinery/pkg/util/wait"

	"kube_web/client"
)

func InitClient() {
	// 定期更新client
	go wait.Forever(client.BuildApiserverClient, 5*time.Hour)
}
