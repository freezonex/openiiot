package tenant

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"freezonex/openiiot/biz/config/consts"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/service/utils/common"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"gorm.io/gen/field"
	"io/ioutil"
	"math/rand"
	"net/http"
)

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

// 测试用
func random() int {
	return rand.Intn(2000) + 30000
}

// AddTenantDB will add tenant record to the DB.
func (a *TenantService) AddTenantDB(ctx context.Context, name string, description string, isDefault bool) (int64, error) {

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": a.s.AuthorizationValue, // 确保 a.s.AuthorizationValue 已定义
	}

	namespace := NewNamespace(name)

	_, err := sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces", a.s.K8SURL), namespace, headers, ctx)

	// 创建 Deployment 结构体的实例
	deployment := NewDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), deployment, headers, ctx)

	webDeployment := WebDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), webDeployment, headers, ctx)

	noderedPersistentVolume := NoderedPersistentVolume(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), noderedPersistentVolume, headers, ctx)

	noderedPersistentVolumeClaim := NoderedPersistentVolumeClaim(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), noderedPersistentVolumeClaim, headers, ctx)

	noderedDeployment := NoderedDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), noderedDeployment, headers, ctx)

	serverDeployment := ServerDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), serverDeployment, headers, ctx)

	grafanaPersistentVolume := GrafanaPersistentVolume(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), grafanaPersistentVolume, headers, ctx)

	grafanaPersistentVolumeClaim := GrafanaPersistentVolumeClaim(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), grafanaPersistentVolumeClaim, headers, ctx)

	grafanaDeployment := GrafanaDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), grafanaDeployment, headers, ctx)

	tdenginePersistentVolumeData := TdenginePersistentVolumeData(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), tdenginePersistentVolumeData, headers, ctx)

	tdenginePersistentVolumeClaimData := TdenginePersistentVolumeClaimData(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), tdenginePersistentVolumeClaimData, headers, ctx)

	tdenginePersistentVolumeLog := TdenginePersistentVolumeLog(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), tdenginePersistentVolumeLog, headers, ctx)

	tdenginePersistentVolumeClaimLog := TdenginePersistentVolumeClaimLog(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), tdenginePersistentVolumeClaimLog, headers, ctx)

	tdenggineDeployment := TdenggineDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), tdenggineDeployment, headers, ctx)

	emqxPersistentVolume := EmqxPersistentVolume(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), emqxPersistentVolume, headers, ctx)

	emqxPersistentVolumeClaim := EmqxPersistentVolumeClaim(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), emqxPersistentVolumeClaim, headers, ctx)

	emqxDeployment := EmqxDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), emqxDeployment, headers, ctx)

	mysqlPersistentVolume := MysqlPersistentVolume(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/persistentvolumes", a.s.K8SURL), mysqlPersistentVolume, headers, ctx)

	mysqlPersistentVolumeClaim := MysqlPersistentVolumeClaim(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/persistentvolumeclaims", a.s.K8SURL, name), mysqlPersistentVolumeClaim, headers, ctx)

	mysqlDeployment := MysqlDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), mysqlDeployment, headers, ctx)

	consolemanagerDeployment := ConsolemanagerDeployment(name)

	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapis/apps/v1/namespaces/%s/deployments", a.s.K8SURL, name), consolemanagerDeployment, headers, ctx)

	//=============================================

	service := NewService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), service, headers, ctx)

	webService := WebService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), webService, headers, ctx)

	noderedService := NoderedService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), noderedService, headers, ctx)

	serverService := ServerService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), serverService, headers, ctx)

	grafanaService := GrafanaService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), grafanaService, headers, ctx)

	tdengineService := TdengineService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), tdengineService, headers, ctx)

	emqxService := EmqxService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), emqxService, headers, ctx)

	mysqlService := MysqlService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), mysqlService, headers, ctx)

	consolemanagerService := ConsolemanagerService(name, random())
	_, err = sendRequest(client, "POST", fmt.Sprintf("%sapi/v1/namespaces/%s/services", a.s.K8SURL, name), consolemanagerService, headers, ctx)

	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)
	existRecord, _ := tx.Where(table.Name.Eq(name)).First()
	if existRecord != nil {
		return -1, errors.New("tenant exist")
	}
	id := common.GetUUID()
	newRecord := &model_openiiot.Tenant{
		ID:          id,
		Name:        name,
		Description: &description,
		IsDefault:   &isDefault,
	}
	err = tx.Create(newRecord)
	if err != nil {
		return -1, err
	}
	return id, nil
}

// GetTenantDB will get tenant record from the DB in condition
func (a *TenantService) GetTenantDB(ctx context.Context, id int64, name string) ([]*model_openiiot.Tenant, error) {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx).Select(field.ALL)
	if id != 0 {
		tx = tx.Where(table.ID.Eq(id))
	}
	if name != "" {
		tx = tx.Where(table.Name.Eq(name))
	}
	tx.Limit(consts.TENANT_RETURN_LIMIT).Order(table.Name)
	data, err := tx.Find()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// UpdateTenantDB will update tenant record from the DB.
func (a *TenantService) UpdateTenantDB(ctx context.Context, id int64, name string, description string) error {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx).Where(table.ID.Eq(id))

	updates := make(map[string]interface{})
	if name != "" {
		updates[table.Name.ColumnName().String()] = name
	}
	if description != "" {
		updates[table.Description.ColumnName().String()] = description
	}

	_, err := tx.Updates(updates)
	return err
}

// DeleteTenantDB will delete tenant record from the DB.
func (a *TenantService) DeleteTenantDB(ctx context.Context, id int64) error {
	table := a.db.DBOpeniiotQuery.Tenant
	tx := table.WithContext(ctx)

	ax := table.WithContext(ctx).Select(field.ALL)
	ax = ax.Where(table.ID.Eq(id))
	data, err := ax.Find()
	if err != nil {
		return err
	}

	// 共享的 HTTP 客户端配置
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	// 设置公共头部
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": a.s.AuthorizationValue, // 确保 a.s.AuthorizationValue 已定义
	}
	_, err = sendRequest(client, "DELETE", fmt.Sprintf("%sapi/v1/namespaces/%s", a.s.K8SURL, data[0].Name), nil, headers, ctx)

	_, err = tx.Where(table.ID.Eq(id)).Delete()

	//kubectl delete pv openiiot-tdengine-volume-data-wenhao
	//kubectl delete pv openiiot-tdengine-volume-log-wenhao
	//kubectl delete pv openiiot-grafana-volume-wenhao
	//kubectl delete pv openiiot-nodered-volume-wenhao
	//kubectl delete pv openiiot-emqx-volume-wenhao
	//kubectl delete pv openiiot-mysql-volume-wenhao

	return err
}
