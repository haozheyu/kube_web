package event

import (
	"kube_web/common"
	"kube_web/controllers/base"
	"kube_web/models"
	"kube_web/resources/event"
	"kube_web/util/logs"
)

type KubeEventController struct {
	base.APIController
}

func (c *KubeEventController) URLMapping() {
	c.Mapping("List", c.List)
}

func (c *KubeEventController) Prepare() {
	// Check administration
	c.APIController.Prepare()

	methodActionMap := map[string]string{
		"List": models.PermissionRead,
	}
	_, method := c.GetControllerAndAction()
	c.PreparePermission(methodActionMap, method, models.PermissionTypeKubePod)
}

func (c *KubeEventController) List() {
	cluster := c.GetString("cluster")
	namespace := c.GetString("namespace", "ALL")
	param := c.BuildKubernetesQueryParam()
	manager := c.Manager(cluster)
	var result *common.Page
	var err error
	result, err = event.GetPodsEventPage(manager.KubeClient, namespace, param)
	if err != nil {
		logs.Error("Get kubernetes events by type error.", cluster, namespace, err)
		c.HandleError(err)
		return
	}

	c.Success(result)
}
