package k8s

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"

	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sService struct {
	db        *mysql.MySQL
	clientset *kubernetes.Clientset
}

// Tag is only used for more than 1 pv and pvc
type K8sUns struct {
	TenantName         string // dt, tenant1, tenant2
	DeploymentCategory string // iot or app, empty string mean iot
	ComponentName      string // nodered, grafana, emqx
	ComponentType      string // frontend, backend, db
	Number             string // 1,2,3,4,5
	Tag                string // now for TDEngine only, data or log
	Alias              string // display name
}

var (
	service *K8sService
	once    sync.Once
)

func NewK8sService(db *mysql.MySQL, k8sConfig *config.K8sConfig) (*K8sService, error) {
	var err error
	once.Do(func() {
		clientset, err := loadKubeConfig()
		if err != nil {
			service = nil
			return
		}

		service = &K8sService{
			db:        db,
			clientset: clientset,
		}
	})
	return service, err
}

func DefaultK8sService() *K8sService {
	return service
}

// loadKubeConfig loads the kubeconfig file from the current directory if it exists,
// otherwise, it falls back to using the InClusterConfig.
func loadKubeConfig() (*kubernetes.Clientset, error) {
	// First, check if a kubeconfig file exists in the current directory
	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	kubeconfig := filepath.Join(currentDir, "test", "localtest", "kubeconfig")
	if _, err := os.Stat(kubeconfig); os.IsNotExist(err) {
		// If the kubeconfig file does not exist in the current directory, use InClusterConfig
		config, err := rest.InClusterConfig()
		if err != nil {
			return nil, err
		}

		// Create the clientset
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			return nil, err
		}

		return clientset, nil
	}

	// If the kubeconfig file exists, use it to build the config
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, err
	}

	// Create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

// int32Ptr is a helper function to create pointers for int32 literals
func int32Ptr(i int32) *int32 {
	return &i
}

// resourceQuantity is a helper function to create resource.Quantity
func resourceQuantity(quantity string) resource.Quantity {
	return resource.MustParse(quantity)
}

// Helper function to create a pointer to a string
func strPtr(s string) *string {
	return &s
}

// Helper function to create a pointer to a PathType
func pathTypePtr(pathType networkingv1.PathType) *networkingv1.PathType {
	return &pathType
}

func sendRequest(client *http.Client, method, url string, data interface{}, headers map[string]string, ctx context.Context) (*http.Response, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, err
	}
	resp.Body.Close()

	logs.CtxInfof(ctx, "%s response status: %s\n%s response body: %s\n", url, resp.Status, url, string(body))
	return resp, nil
}

func (a *K8sService) GetNamespaceName(k8sUns K8sUns) string {
	if k8sUns.TenantName != "" {
		return "openiiot-" + k8sUns.TenantName
	}
	return "openiiot"
}

func (a *K8sService) GetDeploymentName(k8sUns K8sUns) string {
	if k8sUns.DeploymentCategory == "app" {
		return "openiiot-app-" + k8sUns.ComponentName + "-" + k8sUns.ComponentType
	}
	return "openiiot-" + k8sUns.ComponentName + k8sUns.Number
}

func (a *K8sService) GetServicName(k8sUns K8sUns) string {
	if k8sUns.DeploymentCategory == "app" {
		return "openiiot-app-" + k8sUns.ComponentName + "-" + k8sUns.ComponentType + "-service"
	}
	return "openiiot-" + k8sUns.ComponentName + k8sUns.Number + "-service"
}

func (a *K8sService) GetPersistentVolumeName(k8sUns K8sUns) string {
	tagSuffix := ""
	if k8sUns.Tag != "" {
		tagSuffix = "-" + k8sUns.Tag
	}

	if k8sUns.TenantName == "" {
		return "openiiot-" + k8sUns.ComponentName + k8sUns.Number + tagSuffix + "-pv"
	}
	return "openiiot-" + k8sUns.TenantName + "-" + k8sUns.ComponentName + k8sUns.Number + tagSuffix + "-pv"
}

func (a *K8sService) GetPersistentVolumePath(k8sUns K8sUns) string {

	// server mount /volumes to /volumes/openiiot, used to create and delete pv path for others
	if k8sUns.ComponentName == "server" {
		return "/volumes/openiiot"
	}

	tagSuffix := ""
	if k8sUns.Tag != "" {
		tagSuffix = "-" + k8sUns.Tag
	}

	if k8sUns.TenantName == "" {
		return "/volumes/openiiot/" + k8sUns.ComponentName + k8sUns.Number + tagSuffix
	}
	return "/volumes/openiiot/" + k8sUns.TenantName + "/" + k8sUns.ComponentName + k8sUns.Number + tagSuffix
}

// server already mount /volumes to /volumes/openiiot, only used for delete tenant
func (a *K8sService) GetTenantPersistentVolumePath(tenantName string) string {

	return "/volumes/openiiot/" + tenantName
}

func (a *K8sService) GetPersistentVolumeClaimName(k8sUns K8sUns) string {
	tagSuffix := ""
	if k8sUns.Tag != "" {
		tagSuffix = "-" + k8sUns.Tag
	}

	return "openiiot-" + k8sUns.ComponentName + k8sUns.Number + tagSuffix + "-pvc"
}

func (a *K8sService) GetIngressName(k8sUns K8sUns) string {
	if k8sUns.DeploymentCategory == "app" {
		return "openiiot-app-" + k8sUns.ComponentName + "-" + k8sUns.ComponentType + "-ingress"
	}
	return "openiiot-" + k8sUns.ComponentName + k8sUns.Number + "-ingress"
}

func (a *K8sService) GetIngressPath(k8sUns K8sUns) string {
	if k8sUns.ComponentName == "server" {
		return "/api/(.*)"
	}
	return "/" + k8sUns.TenantName + "/" + k8sUns.ComponentName + k8sUns.Number + "/(.*)"
}

func (a *K8sService) GetJobName(k8sUns K8sUns) string {
	return "openiiot-" + k8sUns.ComponentName + k8sUns.Number + "-job"
}

// for grafana only now
func (a *K8sService) GetComponentRootUrl(k8sUns K8sUns) string {

	return "/" + k8sUns.TenantName + "/" + k8sUns.ComponentName + k8sUns.Number
}

func (a *K8sService) GetAppImageName(k8sUns K8sUns) string {
	return "openiiot-app-" + "-" + k8sUns.ComponentName + k8sUns.ComponentType + ":1.0.0"
}
