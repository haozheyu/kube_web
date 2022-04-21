package models_test

import (
	"testing"
	"time"

	"kube_web/models"
)

func TestGetDeployCountByDay(t *testing.T) {
	result, err := models.PublishHistoryModel.GetDeployCountByDay(time.Now().Add(-24*30*time.Hour), time.Now())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
