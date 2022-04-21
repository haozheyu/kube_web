package proxy

import (
	"sort"

	"k8s.io/apimachinery/pkg/util/json"

	"kube_web/client"
	"kube_web/common"
	"kube_web/models/response"
	"kube_web/resources/dataselector"
)

func GetPage(kubeClient client.ResourceHandler, kind string, namespace string, q *common.QueryParam) (*common.Page, error) {
	objs, err := kubeClient.List(kind, namespace, q.LabelSelector)
	if err != nil {
		return nil, err
	}
	commonObjs := make([]dataselector.DataCell, 0)
	for _, obj := range objs {
		objCell, err := getRealObjCellByKind(kind, obj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, objCell)
	}

	sort.Slice(commonObjs, func(i, j int) bool {
		return commonObjs[j].GetProperty(dataselector.NameProperty).Compare(commonObjs[i].GetProperty(dataselector.NameProperty)) == 1
	})

	return dataselector.DataSelectPage(commonObjs, q), nil
}

func GetNames(kubeClient client.ResourceHandler, kind string, namespace string) ([]response.NamesObject, error) {
	objs, err := kubeClient.List(kind, namespace, "")
	if err != nil {
		return nil, err
	}

	commonObjs := make([]response.NamesObject, 0)
	for _, obj := range objs {
		objByte, err := json.Marshal(obj)
		if err != nil {
			return nil, err
		}
		var commonObj ObjectCell
		err = json.Unmarshal(objByte, &commonObj)
		if err != nil {
			return nil, err
		}
		commonObjs = append(commonObjs, response.NamesObject{
			Name: commonObj.Name,
		})
	}

	sort.Slice(commonObjs, func(i, j int) bool {
		return commonObjs[i].Name < commonObjs[j].Name
	})

	return commonObjs, nil
}
