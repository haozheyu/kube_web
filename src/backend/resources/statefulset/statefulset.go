package statefulset

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"kube_web/client"
	"kube_web/client/api"
	"kube_web/resources/common"
	"kube_web/resources/event"
	"kube_web/resources/pod"
	"kube_web/util/maps"
)

type Statefulset struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	Pods       common.PodInfo    `json:"pods"`
}

// GetStatefulsetResource get StatefulSet resource statistics
func GetStatefulsetResource(cli client.ResourceHandler, statefulSet *appsv1.StatefulSet) (*common.ResourceList, error) {
	obj, err := cli.Get(api.ResourceNameStatefulSet, statefulSet.Namespace, statefulSet.Name)
	if err != nil {
		if errors.IsNotFound(err) {
			return common.StatefulsetResourceList(statefulSet), nil
		}
		return nil, err
	}
	old := obj.(*appsv1.StatefulSet)
	oldResourceList := common.StatefulsetResourceList(old)
	newResourceList := common.StatefulsetResourceList(statefulSet)

	return &common.ResourceList{
		Cpu:    newResourceList.Cpu - oldResourceList.Cpu,
		Memory: newResourceList.Memory - oldResourceList.Memory,
	}, nil
}

func CreateOrUpdateStatefulset(cli *kubernetes.Clientset, statefulSet *appsv1.StatefulSet) (*appsv1.StatefulSet, error) {
	old, err := cli.AppsV1().StatefulSets(statefulSet.Namespace).Get(context.TODO(), statefulSet.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().StatefulSets(statefulSet.Namespace).Create(context.TODO(), statefulSet, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = maps.MergeLabels(old.Labels, statefulSet.Labels)
	oldTemplateLabels := old.Spec.Template.Labels
	old.Spec = statefulSet.Spec
	old.Spec.Template.Labels = maps.MergeLabels(oldTemplateLabels, statefulSet.Spec.Template.Labels)

	return cli.AppsV1().StatefulSets(statefulSet.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
}

func GetStatefulsetDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*Statefulset, error) {
	statefulSet, err := cli.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &Statefulset{
		ObjectMeta: common.NewObjectMeta(statefulSet.ObjectMeta),
	}

	podInfo := common.PodInfo{}
	podInfo.Current = statefulSet.Status.ReadyReplicas
	podInfo.Desired = *statefulSet.Spec.Replicas

	pods, err := pod.ListKubePod(indexer, namespace, statefulSet.Spec.Template.Labels)
	if err != nil {
		return nil, err
	}

	podInfo.Warnings, err = event.GetPodsWarningEvents(indexer, pods)
	if err != nil {
		return nil, err
	}

	result.Pods = podInfo

	return result, nil
}
