package k8s

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"io/ioutil"
	"math/rand"
	"net/http"
)

// 测试用
func random() int {
	return rand.Intn(2000) + 30000
}

func K8sServiceCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	//service := NewService(name, random())
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), service, headers, ctx)

	//nginxService := NginxService(name, random())
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), nginxService, headers, ctx)

	webService := WebService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), webService, headers, ctx)

	noderedService := NoderedService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), noderedService, headers, ctx)

	serverService := ServerService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), serverService, headers, ctx)

	grafanaService := GrafanaService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), grafanaService, headers, ctx)

	tdengineService := TdengineService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), tdengineService, headers, ctx)

	emqxService := EmqxService(name, random(), random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), emqxService, headers, ctx)

	mysqlService := MysqlService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), mysqlService, headers, ctx)

	consolemanagerService := ConsolemanagerService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), consolemanagerService, headers, ctx)

	return nil
}

func K8sIngressCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	//noderedIngress := NoderedIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), noderedIngress, headers, ctx)
	//
	//grafanaIngress := GrafanaIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), grafanaIngress, headers, ctx)
	//
	//webIngress := WebIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), webIngress, headers, ctx)

	combinedIngress := CombinedIngress(name)
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), combinedIngress, headers, ctx)

	return nil
}

func K8sJobCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	//noderedIngress := NoderedIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), noderedIngress, headers, ctx)
	//
	//grafanaIngress := GrafanaIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), grafanaIngress, headers, ctx)
	//
	//webIngress := WebIngress(name)
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/networking.k8s.io/v1beta1/namespaces/openiiot-%s/ingresses", K8SURL, name), webIngress, headers, ctx)

	grafanaPluginsJob := CreateInstallGrafanaPluginsJob(name)
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/batch/v1/namespaces/%s/jobs", K8SURL, name), grafanaPluginsJob, headers, ctx)

	return nil
}

func K8sDeploymentPvPvcCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string, id string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	// 创建 Deployment 结构体的实例
	//deployment := NewDeployment(name)
	//
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), deployment, headers, ctx)

	//nginxDeployment := NginxDeployment(name)
	//
	//_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), nginxDeployment, headers, ctx)

	webDeployment := WebDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), webDeployment, headers, ctx)

	noderedPersistentVolume := NoderedPersistentVolume(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), noderedPersistentVolume, headers, ctx)

	noderedPersistentVolumeClaim := NoderedPersistentVolumeClaim(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), noderedPersistentVolumeClaim, headers, ctx)

	noderedDeployment := NoderedDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), noderedDeployment, headers, ctx)

	serverDeployment := ServerDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), serverDeployment, headers, ctx)

	grafanaDataPersistentVolume := GrafanaDataPersistentVolume(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), grafanaDataPersistentVolume, headers, ctx)

	grafanaDataPersistentVolumeClaim := GrafanaDataPersistentVolumeClaim(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), grafanaDataPersistentVolumeClaim, headers, ctx)

	grafanaConfigPersistentVolume := GrafanaConfigPersistentVolume(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), grafanaConfigPersistentVolume, headers, ctx)

	grafanaConfigPersistentVolumeClaim := GrafanaConfigPersistentVolumeClaim(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), grafanaConfigPersistentVolumeClaim, headers, ctx)

	grafanaDeployment := GrafanaDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), grafanaDeployment, headers, ctx)

	tdenginePersistentVolumeData := TdenginePersistentVolumeData(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), tdenginePersistentVolumeData, headers, ctx)

	tdenginePersistentVolumeClaimData := TdenginePersistentVolumeClaimData(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), tdenginePersistentVolumeClaimData, headers, ctx)

	tdenginePersistentVolumeLog := TdenginePersistentVolumeLog(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), tdenginePersistentVolumeLog, headers, ctx)

	tdenginePersistentVolumeClaimLog := TdenginePersistentVolumeClaimLog(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), tdenginePersistentVolumeClaimLog, headers, ctx)

	tdenggineDeployment := TdenggineDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), tdenggineDeployment, headers, ctx)

	emqxPersistentVolume := EmqxPersistentVolume(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), emqxPersistentVolume, headers, ctx)

	emqxPersistentVolumeClaim := EmqxPersistentVolumeClaim(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), emqxPersistentVolumeClaim, headers, ctx)

	emqxDeployment := EmqxDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), emqxDeployment, headers, ctx)

	mysqlPersistentVolume := MysqlPersistentVolume(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), mysqlPersistentVolume, headers, ctx)

	mysqlPersistentVolumeClaim := MysqlPersistentVolumeClaim(name, id)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), mysqlPersistentVolumeClaim, headers, ctx)

	mysqlDeployment := MysqlDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), mysqlDeployment, headers, ctx)

	consolemanagerDeployment := ConsolemanagerDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), consolemanagerDeployment, headers, ctx)

	return nil
}

func K8sConfigmapCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	grafanaConfigmap := GrafanaConfigmap(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/openiiot-%s/configmaps", K8SURL, name), grafanaConfigmap, headers, ctx)

	return nil
}

func K8sNamespaceCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

	namespace := NewNamespace(name)

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces", K8SURL), namespace, headers, ctx)

	return nil

}

func K8sNamespacePvDelete(name string, ctx context.Context, AuthorizationValue string, K8SURL string, id string) *http.Response {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": AuthorizationValue, // 确保 AuthorizationValue 已定义
	}

	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/namespaces/%s", K8SURL, name), nil, headers, ctx)

	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-tdengine-volume-data-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-tdengine-volume-log-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-grafana-data-volume-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-grafana-config-volume-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-nodered-volume-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-emqx-volume-"+name+"-"+id), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-mysql-volume-"+name+"-"+id), nil, headers, ctx)

	return nil
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
