package resource

import (
	"apiserver/pkg/api/apiserver"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

//newTypeMeta create k8s's TypeMeta
func NewTypeMeta(kind, vereion string) metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       kind,
		APIVersion: vereion,
	}
}

func NewObjectMeta(svc *apiserver.Service) metav1.ObjectMeta {
	return metav1.ObjectMeta{
		Name:        svc.Name,
		Namespace:   svc.UserName,
		Labels:      map[string]string{"name": svc.Name},
		Annotations: map[string]string{"name": svc.Name},
	}
}

const (
	ResourceKindConfigMap               = "configmap"
	ResourceKindDaemonSet               = "daemonset"
	ResourceKindDeployment              = "deployment"
	ResourceKindEvent                   = "event"
	ResourceKindHorizontalPodAutoscaler = "horizontalpodautoscaler"
	ResourceKindIngress                 = "ingress"
	ResourceKindJob                     = "job"
	ResourceKindLimitRange              = "limitrange"
	ResourceKindNamespace               = "namespace"
	ResourceKindNode                    = "node"
	ResourceKindPersistentVolumeClaim   = "persistentvolumeclaim"
	ResourceKindPersistentVolume        = "persistentvolume"
	ResourceKindPod                     = "pod"
	ResourceKindReplicaSet              = "replicaset"
	ResourceKindReplicationController   = "replicationcontroller"
	ResourceKindResourceQuota           = "resourcequota"
	ResourceKindSecret                  = "secret"
	ResourceKindService                 = "service"
	ResourceKindStatefulSet             = "statefulset"
	ResourceKindThirdPartyResource      = "thirdpartyresource"
	ResourceKindStorageClass            = "storageclass"
	ResourceKindRbacRole                = "role"
	ResourceKindRbacClusterRole         = "clusterrole"
	ResourceKindRbacRoleBinding         = "rolebinding"
	ResourceKindRbacClusterRoleBinding  = "clusterrolebinding"
)

// type AppStatus int32
// type UpdateStatus int32

const (
	AppBuilding  = 0
	AppSuccessed = 1
	AppFailed    = 2
	AppRunning   = 3
	AppStop      = 4
	AppDelete    = 5
	AppUnknow    = 6

	StartFailed    = 10
	StartSuccessed = 11

	StopFailed    = 20
	StopSuccessed = 21

	ScaleFailed    = 30
	ScaleSuccessed = 31

	UpdateConfigFailed    = 40
	UpdateConfigSuccessed = 41

	RedeploymentFailed    = 50
	RedeploymentSuccessed = 51
)

var (
	Status = map[int]string{
		0: "AppBuilding",
		1: "AppSuccessed",
		2: "AppFailed",
		3: "AppRunning",
		4: "AppStop",
		5: "AppDelete",
		6: "AppUnknow",
	}

	ListEverything = metav1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	}
)
