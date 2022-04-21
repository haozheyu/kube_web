package models_test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"kube_web/common"
	"kube_web/models"
)

func TestGetAll(t *testing.T) {

	fmt.Println(models.GetTotal(models.TableNameDeployment, &common.QueryParam{
		Query: map[string]interface{}{
			"app__id": 3,
		},
	}))

}
