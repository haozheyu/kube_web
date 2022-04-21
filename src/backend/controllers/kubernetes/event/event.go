package event

import (
	"fmt"
	"net/http"

	"kube_web/client/api"
	"kube_web/common"
	"kube_web/controllers/base"
	"kube_web/models"
	erroresult "kube_web/models/response/errors"
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

// @Title GetPodEvent
// @Description Get Pod Event by resource type and name
// @Param	pageNo		query 	int	false		"the page current no"
// @Param	pageSize		query 	int	false		"the page size"
// @Param	type		query 	string	true		"the query type. deployments, statefulsets, daemonsets,cronjobs"
// @Param	name		query 	string	true		"the query resource name."
// @Success 200 {object} models.Deployment success
// @router /namespaces/:namespace/clusters/:cluster [get]
func (c *KubeEventController) List() {
	cluster := c.Ctx.Input.Param(":cluster")
	namespace := c.Ctx.Input.Param(":namespace")
	resourceType := c.Ctx.Input.Query("type")
	resourceName := c.Ctx.Input.Query("name")
	param := c.BuildKubernetesQueryParam()
	manager := c.Manager(cluster)
	var result *common.Page
	var err error
	switch resourceType {
	case api.ResourceNameCronJob:
		result, err = event.GetPodsEventByCronJobPage(manager.KubeClient, namespace, resourceName, param)
	default:
		err = &erroresult.ErrorResult{
			Code: http.StatusBadRequest,
			Msg:  fmt.Sprintf("Unsupported resource type (%s). ", resourceType),
		}
	}
	if err != nil {
		logs.Error("Get kubernetes events by type error.", cluster, namespace, resourceType, resourceName, err)
		c.HandleError(err)
		return
	}
	c.Success(result)
}
