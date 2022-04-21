package proxy

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"kube_web/resources/common"
	"kube_web/resources/dataselector"
)

type ObjectCell common.Object

func (cell ObjectCell) GetProperty(name dataselector.PropertyName) dataselector.ComparableValue {
	return baseProperty(name, cell.ObjectMeta)
}

func baseProperty(name dataselector.PropertyName, meta metav1.ObjectMeta) dataselector.ComparableValue {
	switch name {
	case dataselector.NameProperty:
		return dataselector.StdComparableString(meta.Name)
	case dataselector.CreationTimestampProperty:
		return dataselector.StdComparableTime(meta.CreationTimestamp.Time)
	case dataselector.NamespaceProperty:
		return dataselector.StdComparableString(meta.Namespace)

	case dataselector.ReferenceUIDProperty:
		refs := meta.OwnerReferences
		for i := range refs {
			if refs[i].Controller != nil && *refs[i].Controller {
				return dataselector.StdComparableString(refs[i].UID)
			}
		}
		return nil

	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}

type PodCell v1.Pod

func (cell PodCell) GetProperty(name dataselector.PropertyName) dataselector.ComparableValue {
	switch name {
	case dataselector.PodIPProperty:
		return dataselector.StdComparableString(cell.Status.PodIP)
	case dataselector.NodeNameProperty:
		return dataselector.StdComparableString(cell.Spec.NodeName)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return baseProperty(name, cell.ObjectMeta)
	}
}
