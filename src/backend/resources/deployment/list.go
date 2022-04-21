package deployment

import (
	"kube_web/client"
	"kube_web/common"
	"kube_web/resources/dataselector"
)

func GetDeploymentPage(indexer *client.CacheFactory, namespace string, q *common.QueryParam) (*common.Page, error) {
	kubeDeployments, err := GetDeploymentList(indexer, namespace)
	if err != nil {
		return nil, err
	}

	deployments := make([]Deployment, 0)

	for i := 0; i < len(kubeDeployments); i++ {
		deploy, err := toDeployment(kubeDeployments[i], indexer)
		if err != nil {
			return nil, err
		}
		deployments = append(deployments, *deploy)
	}

	return dataselector.DataSelectPage(toCells(deployments), q), nil
}

func toCells(deploy []Deployment) []dataselector.DataCell {
	cells := make([]dataselector.DataCell, len(deploy))
	for i := range deploy {
		cells[i] = DeploymentCell(deploy[i])
	}
	return cells
}
