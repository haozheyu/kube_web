package pvc

import (
	"context"
	"errors"
	"strings"

	kapi "k8s.io/api/core/v1"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"kube_web/resources/pv"
)

var (
	RbdNotFoundError = errors.New("image not found")
)

func CreateOrUpdatePersistentVolumeClaim(cli *kubernetes.Clientset, pvc *kapi.PersistentVolumeClaim) (*kapi.PersistentVolumeClaim, error) {
	old, err := cli.CoreV1().PersistentVolumeClaims(pvc.Namespace).Get(context.TODO(), pvc.Name, metaV1.GetOptions{})
	if err != nil {
		if kerrors.IsNotFound(err) {
			return cli.CoreV1().PersistentVolumeClaims(pvc.Namespace).Create(context.TODO(), pvc, metaV1.CreateOptions{})
		}
		return nil, err
	}
	old.Labels = pvc.Labels
	old.Annotations = pvc.Annotations

	return cli.CoreV1().PersistentVolumeClaims(pvc.Namespace).Update(context.TODO(), old, metaV1.UpdateOptions{})
}

func GetPersistentVolumeClaimDetail(cli *kubernetes.Clientset, name, namespace string) (*kapi.PersistentVolumeClaim, error) {
	return cli.CoreV1().
		PersistentVolumeClaims(namespace).
		Get(context.TODO(), name, metaV1.GetOptions{})
}

func DeletePersistentVolumeClaim(cli *kubernetes.Clientset, name, namespace string) error {
	return cli.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), name, metaV1.DeleteOptions{})
}

func GetRbdImageByPvc(cli *kubernetes.Clientset, name, namespace string) (string, error) {
	pvcResult, err := GetPersistentVolumeClaimDetail(cli, name, namespace)
	if err != nil {
		return "", err
	}

	if pvcResult.Spec.VolumeName != "" {
		pvResult, err := pv.GetPersistentVolumeByName(cli, pvcResult.Spec.VolumeName)
		if err != nil {
			return "", err
		}
		// 获取RBD状态
		if pvResult.Spec.RBD != nil {
			return pvResult.Spec.RBD.RBDImage, nil
		} else if pvResult.Spec.CephFS != nil {
			paths := strings.Split(pvResult.Spec.CephFS.Path, "/")
			return paths[len(paths)-1], nil
		}
	}
	return "", RbdNotFoundError
}

func GetImageNameAndTypeByPvc(cli *kubernetes.Clientset, name, namespace string) (string, string, error) {
	pvcResult, err := GetPersistentVolumeClaimDetail(cli, name, namespace)
	if err != nil {
		return "", "", err
	}

	if pvcResult.Spec.VolumeName != "" {
		pvResult, err := pv.GetPersistentVolumeByName(cli, pvcResult.Spec.VolumeName)
		if err != nil {
			return "", "", err
		}
		// 获取RBD状态
		if pvResult.Spec.RBD != nil {
			return pvResult.Spec.RBD.RBDImage, "rbd", nil
		}
		if pvResult.Spec.CephFS != nil {
			paths := strings.Split(pvResult.Spec.CephFS.Path, "/")
			return paths[len(paths)-1], "cephfs", nil
		}
		if pvResult.Spec.Glusterfs != nil {
			paths := strings.Split(pvResult.Spec.Glusterfs.Path, "/")
			return paths[len(paths)-1], "glusterfs", nil
		}
		if pvResult.Spec.NFS != nil {
			paths := strings.Split(pvResult.Spec.NFS.Path, "/")
			return paths[len(paths)-1], "nfs", nil
		}
	}
	return "", "", RbdNotFoundError
}
