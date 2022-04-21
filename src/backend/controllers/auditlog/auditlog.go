package auditlog

import (
	"kube_web/controllers/base"
	"kube_web/models"
	"kube_web/util/logs"
)

type AuditLogController struct {
	base.APIController
}

func (c *AuditLogController) URLMapping() {
	c.Mapping("List", c.List)
}

func (c *AuditLogController) Prepare() {
	// Check administration
	c.APIController.Prepare()
}

// @Title GetAll
// @Description get audit log
// @Param	pageNo		query 	int	false		"the page current no"
// @Param	pageSize		query 	int	false		"the page size"
// @Success 200 {object} []models.AuditLog success
// @router / [get]
func (c *AuditLogController) List() {
	param := c.BuildQueryParam()
	user := c.Ctx.Input.Query("user")
	if user != "" {
		param.Query["User__contains"] = user
	}

	userIp := c.Ctx.Input.Query("userIp")
	if userIp != "" {
		param.Query["UserIp__contains"] = userIp
	}

	logType := c.Ctx.Input.Query("logType")
	if logType != "" {
		param.Query["LogType__contains"] = logType
	}

	logLevel := c.Ctx.Input.Query("logLevel")
	if logLevel != "" {
		param.Query["LogLevel__contains"] = logLevel
	}

	action := c.Ctx.Input.Query("action")
	if action != "" {
		param.Query["Action__contains"] = action
	}

	auditLogs := []models.AuditLog{}

	total, err := models.GetTotal(new(models.AuditLog), param)
	if err != nil {
		logs.Error("get total count by param (%v) error. %v", param, err)
		c.HandleError(err)
		return
	}

	err = models.GetAll(new(models.AuditLog), &auditLogs, param)
	if err != nil {
		logs.Error("list by param (%v) error. %v", param, err)
		c.HandleError(err)
		return
	}

	c.Success(param.NewPage(total, auditLogs))
	return
}
