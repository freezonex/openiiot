package tenant

import (
	"freezonex/openiiot/biz/config"
	"sync"

	"freezonex/openiiot/biz/dal/mysql"
)

type TenantService struct {
	db *mysql.MySQL
	s  *config.K8sConfig
}

// namespace
type Metadata struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

// 定义主结构体
type Namespace struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
}

// Deployment 定义了 Kubernetes Deployment 的结构
type Deployment struct {
	APIVersion string   `json:"apiVersion"`
	Kind       string   `json:"kind"`
	Metadata   Metadata `json:"metadata"`
	Spec       Spec     `json:"spec"`
}

type EnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Spec 定义了 Deployment 的规格
type Spec struct {
	Replicas int         `json:"replicas"`
	Selector Selector    `json:"selector"`
	Template PodTemplate `json:"template"`
}

// Selector 定义了如何选择 Pod
type Selector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}

// PodTemplate 定义了 Pod 的模板
type PodTemplate struct {
	Metadata Metadata `json:"metadata"`
	Spec     PodSpec  `json:"spec"`
}

// PodSpec 定义了 Pod 的规格
type PodSpec struct {
	Containers []Container `json:"containers"`
	Volumes    []Volume    `json:"volumes,omitempty"`
}

// Container 定义了 Pod 中的容器
type Container struct {
	Image        string               `json:"image"`
	Name         string               `json:"name"`
	Resources    ResourceRequirements `json:"resources"`
	VolumeMounts []VolumeMount        `json:"volumeMounts,omitempty"`
	Env          []EnvVar             `json:"env"`
}

type VolumeMount struct {
	Name      string `json:"name"`
	MountPath string `json:"mountPath"`
}

type Volume struct {
	Name                  string                `json:"name"`
	PersistentVolumeClaim PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

type PersistentVolumeClaim struct {
	ClaimName  string                    `json:"claimName"`
	APIVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   Metadata                  `json:"metadata"`
	Spec       PersistentVolumeClaimSpec `json:"spec"`
}

// Service 结构体代表 Kubernetes Service 资源
type Service struct {
	APIVersion string      `json:"apiVersion"`
	Kind       string      `json:"kind"`
	Metadata   Metadata    `json:"metadata"`
	Spec       ServiceSpec `json:"spec"`
}

// ServiceSpec 结构体用于描述 Service 的规格
type ServiceSpec struct {
	Ports    []Port            `json:"ports"`
	Selector map[string]string `json:"selector"`
	Type     string            `json:"type"`
}

// Port 结构体定义了 Service 的端口配置
type Port struct {
	Port       int    `json:"port"`
	Protocol   string `json:"protocol"`
	TargetPort int    `json:"targetPort"`
	NodePort   int    `json:"nodePort"`
}

type ResourceRequirements struct {
	Requests map[string]string `json:"requests"`
	Limits   map[string]string `json:"limits"`
}

type PersistentVolumeClaimSpec struct {
	AccessModes []string `json:"accessModes"`
	Resources   struct {
		Requests struct {
			Storage string `json:"storage"`
		} `json:"requests"`
	} `json:"resources"`
	VolumeName       string `json:"volumeName,omitempty"`
	StorageClassName string `json:"storageClassName,omitempty"` // 添加这行
}

type PersistentVolumeSpec struct {
	Capacity                      map[string]string `json:"capacity"`
	AccessModes                   []string          `json:"accessModes"`
	PersistentVolumeReclaimPolicy string            `json:"persistentVolumeReclaimPolicy"`
	StorageClassName              string            `json:"storageClassName"`
	HostPath                      struct {
		Path string `json:"path"`
	} `json:"hostPath"`
	// Add other relevant fields based on your requirements
}

type PersistentVolume struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Spec PersistentVolumeSpec `json:"spec"`
}

var (
	service *TenantService
	once    sync.Once
)

func NewTenantService(db *mysql.MySQL, s *config.K8sConfig) *TenantService {
	once.Do(func() {
		service = &TenantService{
			db: db,
			s:  s,
		}
	})
	return service
}
