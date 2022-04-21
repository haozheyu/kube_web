package service

import (
	"context"
	kapi "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CreateOrUpdateService(cli *kubernetes.Clientset, service *kapi.Service) (*kapi.Service, error) {
	old, err := cli.CoreV1().Services(service.Namespace).Get(context.TODO(), service.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return cli.CoreV1().Services(service.Namespace).Create(context.TODO(), service, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = service.Labels
	old.Spec.ExternalIPs = service.Spec.ExternalIPs
	old.Spec.Selector = service.Spec.Selector
	old.Spec.Ports = service.Spec.Ports

	return cli.CoreV1().Services(service.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
}
