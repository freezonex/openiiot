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

	emqxService := EmqxService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), emqxService, headers, ctx)

	mysqlService := MysqlService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), mysqlService, headers, ctx)

	consolemanagerService := ConsolemanagerService(name, random())
	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", K8SURL, name), consolemanagerService, headers, ctx)

	return nil
}

func K8sDeploymentPvPvcCreate(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

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

	webDeployment := WebDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), webDeployment, headers, ctx)

	noderedPersistentVolume := NoderedPersistentVolume(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), noderedPersistentVolume, headers, ctx)

	noderedPersistentVolumeClaim := NoderedPersistentVolumeClaim(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), noderedPersistentVolumeClaim, headers, ctx)

	noderedDeployment := NoderedDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), noderedDeployment, headers, ctx)

	serverDeployment := ServerDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), serverDeployment, headers, ctx)

	grafanaPersistentVolume := GrafanaPersistentVolume(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), grafanaPersistentVolume, headers, ctx)

	grafanaPersistentVolumeClaim := GrafanaPersistentVolumeClaim(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), grafanaPersistentVolumeClaim, headers, ctx)

	grafanaDeployment := GrafanaDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), grafanaDeployment, headers, ctx)

	tdenginePersistentVolumeData := TdenginePersistentVolumeData(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), tdenginePersistentVolumeData, headers, ctx)

	tdenginePersistentVolumeClaimData := TdenginePersistentVolumeClaimData(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), tdenginePersistentVolumeClaimData, headers, ctx)

	tdenginePersistentVolumeLog := TdenginePersistentVolumeLog(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), tdenginePersistentVolumeLog, headers, ctx)

	tdenginePersistentVolumeClaimLog := TdenginePersistentVolumeClaimLog(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), tdenginePersistentVolumeClaimLog, headers, ctx)

	tdenggineDeployment := TdenggineDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), tdenggineDeployment, headers, ctx)

	emqxPersistentVolume := EmqxPersistentVolume(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), emqxPersistentVolume, headers, ctx)

	emqxPersistentVolumeClaim := EmqxPersistentVolumeClaim(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), emqxPersistentVolumeClaim, headers, ctx)

	emqxDeployment := EmqxDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), emqxDeployment, headers, ctx)

	mysqlPersistentVolume := MysqlPersistentVolume(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", K8SURL), mysqlPersistentVolume, headers, ctx)

	mysqlPersistentVolumeClaim := MysqlPersistentVolumeClaim(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", K8SURL, name), mysqlPersistentVolumeClaim, headers, ctx)

	mysqlDeployment := MysqlDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), mysqlDeployment, headers, ctx)

	consolemanagerDeployment := ConsolemanagerDeployment(name)

	_, _ = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", K8SURL, name), consolemanagerDeployment, headers, ctx)

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

func K8sNamespacePvDelete(name string, ctx context.Context, AuthorizationValue string, K8SURL string) *http.Response {

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

	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-tdengine-volume-data-"+name), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-tdengine-volume-log-"+name), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-grafana-volume-"+name), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-nodered-volume-"+name), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-emqx-volume-"+name), nil, headers, ctx)
	_, _ = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/persistentvolumes/%s", K8SURL, "openiiot-mysql-volume-"+name), nil, headers, ctx)

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
