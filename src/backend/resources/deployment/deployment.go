package deployment

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"

	"kube_web/client"
	"kube_web/client/api"
	erroresult "kube_web/models/response/errors"
	"kube_web/resources/common"
	"kube_web/resources/event"
	"kube_web/resources/pod"
	"kube_web/util/maps"
)

type Deployment struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	Pods       common.PodInfo    `json:"pods"`
	Containers []string          `json:"containers"`
}

func GetDeploymentList(indexer *client.CacheFactory, namespace string) ([]*appsv1.Deployment, error) {
	deployments, err := indexer.DeploymentLister().Deployments(namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	sort.Slice(deployments, func(i, j int) bool {
		return deployments[i].Name > deployments[j].Name
	})
	return deployments, nil
}

// GetDeploymentResource get deployment resource statistics
func GetDeploymentResource(cli client.ResourceHandler, deployment *appsv1.Deployment) (*common.ResourceList, error) {
	obj, err := cli.Get(api.ResourceNameDeployment, deployment.Namespace, deployment.Name)

	if err != nil {
		if errors.IsNotFound(err) {
			return common.DeploymentResourceList(deployment), nil
		}
		return nil, err
	}
	old := obj.(*appsv1.Deployment)
	oldResourceList := common.DeploymentResourceList(old)
	newResourceList := common.DeploymentResourceList(deployment)

	return &common.ResourceList{
		Cpu:    newResourceList.Cpu - oldResourceList.Cpu,
		Memory: newResourceList.Memory - oldResourceList.Memory,
	}, nil
}

func CreateOrUpdateDeployment(cli *kubernetes.Clientset, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	//old, err := cli.AppsV1beta1().Deployments(deployment.Namespace).Get(deployment.Name, metaV1.GetOptions{})
	old, err := cli.AppsV1().Deployments(deployment.Namespace).Get(context.TODO(), deployment.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.AppsV1().Deployments(deployment.Namespace).Create(context.TODO(), deployment, metaV1.CreateOptions{})
		}
		return nil, err
	}
	err = checkDeploymentLabelSelector(deployment, old)
	if err != nil {
		return nil, err
	}

	old.Labels = deployment.Labels
	old.Annotations = deployment.Annotations
	old.Spec = deployment.Spec

	return cli.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
}

func UpdateDeployment(cli *kubernetes.Clientset, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	old, err := cli.AppsV1().Deployments(deployment.Namespace).Get(context.TODO(), deployment.Name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	err = checkDeploymentLabelSelector(deployment, old)
	if err != nil {
		return nil, err
	}

	return cli.AppsV1().Deployments(deployment.Namespace).Update(context.TODO(), deployment, metaV1.UpdateOptions{})
}

// check Deployment .Spec.Selector.MatchLabels, prevent orphan ReplicaSet
// old deployment .Spec.Selector.MatchLabels labels should contain all new deployment .Spec.Selector.MatchLabels labels
// e.g. old Deployment .Spec.Selector.MatchLabels is app = infra-wayne,wayne-app = infra
// new Deployment .Spec.Selector.MatchLabels valid labels is
// app = infra-wayne or wayne-app = infra or app = infra-wayne,wayne-app = infra
func checkDeploymentLabelSelector(new *appsv1.Deployment, old *appsv1.Deployment) error {
	for key, value := range new.Spec.Selector.MatchLabels {
		oldValue, ok := old.Spec.Selector.MatchLabels[key]
		if !ok || oldValue != value {
			return &erroresult.ErrorResult{
				Code: http.StatusBadRequest,
				Msg: fmt.Sprintf("New's Deployment MatchLabels(%s) not match old MatchLabels(%s), do not allow deploy to prevent the orphan ReplicaSet. ",
					maps.LabelsToString(new.Spec.Selector.MatchLabels), maps.LabelsToString(old.Spec.Selector.MatchLabels)),
			}
		}
	}

	return nil
}

func GetDeployment(cli *kubernetes.Clientset, name, namespace string) (*appsv1.Deployment, error) {
	return cli.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
}

func GetDeploymentDetail(cli *kubernetes.Clientset, indexer *client.CacheFactory, name, namespace string) (*Deployment, error) {
	deployment, err := cli.AppsV1().Deployments(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return toDeployment(deployment, indexer)
}

func toDeployment(deployment *appsv1.Deployment, indexer *client.CacheFactory) (*Deployment, error) {
	result := &Deployment{
		ObjectMeta: common.NewObjectMeta(deployment.ObjectMeta),
	}

	podInfo := common.PodInfo{}
	podInfo.Current = deployment.Status.AvailableReplicas
	podInfo.Desired = *deployment.Spec.Replicas
	var err error

	pods, err := pod.ListKubePod(indexer, deployment.Namespace, deployment.Spec.Template.Labels)
	if err != nil {
		return nil, err
	}

	podInfo.Warnings, err = event.GetPodsWarningEvents(indexer, pods)
	if err != nil {
		return nil, err
	}

	result.Pods = podInfo

	containers := make([]string, 0)
	for _, container := range deployment.Spec.Template.Spec.Containers {
		containers = append(containers, container.Image)
	}

	result.Containers = containers

	return result, nil
}

func DeleteDeployment(cli *kubernetes.Clientset, name, namespace string) error {
	deletionPropagation := metaV1.DeletePropagationBackground
	return cli.AppsV1().
		Deployments(namespace).
		Delete(context.TODO(), name, metaV1.DeleteOptions{PropagationPolicy: &deletionPropagation})
}

func UpdateScale(cli *kubernetes.Clientset, deploymentname string, namespace string, newreplica int32) error {
	deployments := cli.AppsV1beta1().Deployments(namespace)
	deployment, err := deployments.Get(context.TODO(), deploymentname, metaV1.GetOptions{})
	if err != nil {
		return err
	}
	deployment.Spec.Replicas = &newreplica
	_, err = deployments.Update(context.TODO(), deployment, metaV1.UpdateOptions{})
	return err
}
