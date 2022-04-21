package models_test

import (
	"testing"

	"kube_web/models"
)

func TestUpdateOrders(t *testing.T) {
	deployments := []*models.Deployment{
		{
			Id:      1,
			OrderId: 2,
		},
		{
			Id:      2,
			OrderId: 1,
		},
	}
	err := models.DeploymentModel.UpdateOrders(deployments)

	if err != nil {
		t.Errorf("update orders error.%v ", err)
	}
}
