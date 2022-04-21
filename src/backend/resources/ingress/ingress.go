package ingress

import (
	"context"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"kube_web/util/logs"
)

func CreateOrUpdateIngress(c *kubernetes.Clientset, ingress *v1beta1.Ingress) (*Ingress, error) {
	old, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Get(context.TODO(), ingress.Name, metaV1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			kubeIngress, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Create(context.TODO(), ingress, metaV1.CreateOptions{})
			if err != nil {
				return nil, err
			}
			return toIngress(kubeIngress), nil
		}
		return nil, err
	}

	// ingress.Spec.DeepCopyInto(&old.Spec)
	// also need update Labels、Annotations、Spec
	old.Labels = ingress.Labels
	old.Annotations = ingress.Annotations
	old.Spec = ingress.Spec
	logs.Info("new ingress", old)

	kubeIngress, err := c.ExtensionsV1beta1().Ingresses(ingress.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return toIngress(kubeIngress), nil
}

func GetIngressDetail(c *kubernetes.Clientset, name, namespace string) (*Ingress, error) {
	ingress, err := c.ExtensionsV1beta1().Ingresses(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return toIngress(ingress), nil
}

func GetIngress(c *kubernetes.Clientset, name, namespace string) (ingress *v1beta1.Ingress, err error) {
	ingress, err = c.ExtensionsV1beta1().Ingresses(namespace).Get(context.TODO(), name, metaV1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return
}

func DeleteIngress(c *kubernetes.Clientset, name, namespace string) error {
	return c.ExtensionsV1beta1().Ingresses(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func GetIngressList(cli *kubernetes.Clientset, namespace string, opts metaV1.ListOptions) (list []*Ingress, err error) {
	kubeIngressList, err := cli.ExtensionsV1beta1().Ingresses(namespace).List(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	for _, kubeIngress := range kubeIngressList.Items {
		list = append(list, toIngress(&kubeIngress))
	}
	return
}
