package crd

import (
	"context"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"

	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube_web/common"
	"kube_web/resources/dataselector"
)

func GetCRDPage(cli *clientset.Clientset, q *common.QueryParam) (*common.Page, error) {
	crdList, err := cli.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metaV1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return dataselector.DataSelectPage(toCells(crdList.Items), q), nil
}

func toCells(deploy []apiextensions.CustomResourceDefinition) []dataselector.DataCell {
	cells := make([]dataselector.DataCell, len(deploy))
	for i := range deploy {
		cells[i] = CRDCell(deploy[i])
	}
	return cells
}
