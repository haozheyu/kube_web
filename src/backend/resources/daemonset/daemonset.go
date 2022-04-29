package daemonset

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"kube_web/client"
	"kube_web/resources/common"
	"kube_web/resources/event"
	"kube_web/resources/pod"
	"kube_web/util/maps"
)

type DaemonSet struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	Pods       common.PodInfo    `json:"pods"`
}

func CreateOrUpdateDaemonSet(cli *kubernetes.Clientset, daemonSet *appsv1.DaemonSet) (*appsv1.DaemonSet, error) {
	old, err := cli.AppsV1().DaemonSets(daemonSet.Namespace).Get(context.TODO(), daemonSet.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().DaemonSets(daemonSet.Namespace).Create(context.TODO(), daemonSet, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = maps.MergeLabels(old.Labels, daemonSet.Labels)
	oldTemplateLabels := old.Spec.Template.Labels
	old.Spec = daemonSet.Spec
	old.Spec.Template.Labels = maps.MergeLabels(oldTemplateLabels, daemonSet.Spec.Template.Labels)

	return cli.AppsV1().DaemonSets(daemonSet.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
}

func GetDaemonSetDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*DaemonSet, error) {
	daemonSet, err := cli.AppsV1().DaemonSets(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	result := &DaemonSet{
		ObjectMeta: common.NewObjectMeta(daemonSet.ObjectMeta),
	}

	podInfo := common.PodInfo{}
	podInfo.Current = daemonSet.Status.NumberAvailable
	podInfo.Desired = daemonSet.Status.DesiredNumberScheduled

	pods, err := pod.ListKubePod(indexer, namespace, daemonSet.Spec.Template.Labels)
	if err != nil {
		return nil, err
	}

	podInfo.Warnings, err = event.GetPodsEvents(indexer, pods)
	if err != nil {
		return nil, err
	}

	result.Pods = podInfo

	return result, nil
}

func DeleteDaemonSet(cli *kubernetes.Clientset, name, namespace string) error {
	deletionPropagation := metaV1.DeletePropagationBackground
	return cli.AppsV1().
		DaemonSets(namespace).
		Delete(context.TODO(), name, metaV1.DeleteOptions{PropagationPolicy: &deletionPropagation})
}
