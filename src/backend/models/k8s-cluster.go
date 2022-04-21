package models

import (
	"github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/client/orm"
)

type K8SCluster struct {
	BaseModel
	ClusterName    string `json:"clusterName" gorm:"comment:集群名称" form:"clusterName" binding:"required"`
	KubeConfig     string `json:"kubeConfig" gorm:"comment:集群凭证;type:varchar(12800)" form:"kubeConfig" binding:"required"`
	ClusterVersion string `json:"clusterVersion" gorm:"comment:集群版本"`
	NodeNumber     int    `json:"nodeNumber" gorm:"comment:节点数"`
}

func (ks K8SCluster) TableName() string {
	var k BaseModel
	return k.TableName("k8s_cluster")
}

type ClusterVersion struct {
	BaseModel
	Version string `json:"version"`
}

func (v ClusterVersion) TableName() string {
	var k BaseModel
	return k.TableName("k8s_cluster_version")
}

type PaginationQ struct {
	Size    int    `form:"size" json:"size"`
	Page    int    `form:"page" json:"page"`
	Total   int64  `json:"total"`
	Keyword string `form:"keyword" json:"keyword"`
}

type ClusterIds struct {
	Data interface{} `json:"clusterIds"`
}

type ClusterNodesStatus struct {
	BaseModel
	NodeCount       int     `json:"node_count"`
	Ready           int     `json:"ready"`
	UnReady         int     `json:"unready"`
	Namespace       int     `json:"namespace"`
	Deployment      int     `json:"deployment"`
	Pod             int     `json:"pod"`
	CpuUsage        float64 `json:"cpu_usage" desc:"cpu使用率"`
	CpuCore         float64 `json:"cpu_core"`
	CpuCapacityCore float64 `json:"cpu_capacity_core"`
	MemoryUsage     float64 `json:"memory_usage" desc:"内存使用率"`
	MemoryUsed      float64 `json:"memory_used"`
	MemoryTotal     float64 `json:"memory_total"`
}

func init() {
	orm.RegisterModel(new(K8SCluster), new(ClusterVersion))
}

func CreateK8SCluster(cluster K8SCluster) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&cluster)
	return err
}

func ListK8SCluster(currentpage int, pagesize int, conditions string) ([]K8SCluster, int64, error) {
	var ks K8SCluster
	var kss []K8SCluster
	totalItem, totalpages, seter := GetPagesInfo(ks.TableName(), currentpage, pagesize, conditions)
	_, err := seter.QueryRows(&kss)
	logs.Info("获取添加集群信息列表", totalItem, totalpages)
	return kss, int64(totalpages), err
}

func GetK8sCluster(id uint) (K8sCluster K8SCluster, err error) {
	o := orm.NewOrm()
	var cluter K8SCluster
	err = o.QueryTable(K8sCluster.TableName()).Filter("id", id).One(&cluter)
	if err != nil {
		return cluter, err
	}
	return cluter, nil
}

func DelCluster(id uint) (err error) {
	var k K8SCluster
	o := orm.NewOrm()
	_, err = o.QueryTable(k.TableName()).Filter("id", id).Delete()
	return err
}
