package namespace

import (
	"k8s.io/api/core/v1"

	"kube_web/resources/common"
	"kube_web/resources/dataselector"
)

type Namespace struct {
	ObjectMeta common.ObjectMeta `json:"objectMeta"`
	Status     v1.NamespacePhase `json:"status"`
}

func toNamespace(namespace *v1.Namespace) *Namespace {
	result := &Namespace{
		ObjectMeta: common.NewObjectMeta(namespace.ObjectMeta),
	}
	result.Status = namespace.Status.Phase

	return result
}

type NamespaceCell Namespace

func (cell NamespaceCell) GetProperty(name dataselector.PropertyName) dataselector.ComparableValue {
	switch name {
	case dataselector.NameProperty:
		return dataselector.StdComparableString(cell.ObjectMeta.Name)
	case dataselector.CreationTimestampProperty:
		return dataselector.StdComparableTime(cell.ObjectMeta.CreationTimestamp.Time)
	case dataselector.NamespaceProperty:
		return dataselector.StdComparableString(cell.ObjectMeta.Namespace)
	default:
		// if name is not supported then just return a constant dummy value, sort will have no effect.
		return nil
	}
}
