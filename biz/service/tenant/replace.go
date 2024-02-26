package tenant

import "freezonex/openiiot/biz/model/freezonex_openiiot_api"

func emqxreplace1(name string) *freezonex_openiiot_api.BridgeConfig {
	return &freezonex_openiiot_api.BridgeConfig{
		Name:             "air_quality",
		Type:             "webhook",
		Body:             "INSERT INTO demo.${payload.tableName} USING demo.air_quality TAGS('${payload.nodeId}', '${payload.area}') VALUES(now,${payload.temperature},${payload.humidity},${payload.volume},${payload.pm10},${payload.pm25},${payload.so2},${payload.no2},${payload.co},'${payload.status}')",
		ConnectTimeout:   "15s",
		Enable:           true,
		EnablePipelining: 100,
		Headers: map[string]string{
			"accept":        "application/json",
			"cache-control": "no-cache",
			"connection":    "keep-alive",
			"content-type":  "application/json",
			"keep-alive":    "timeout=5",
			"Authorization": "Basic cm9vdDp0YW9zZGF0YQ==",
		},
		LocalTopic: "emqx_http/#",
		MaxRetries: 3,
		Method:     "post",
		PoolSize:   4,
		PoolType:   "random",
		ResourceOpts: &freezonex_openiiot_api.ResourceOpts{ // Instantiating and assigning a pointer
			WorkerPoolSize:      1,
			MaxBufferBytes:      104857600,
			InflightWindow:      100,
			QueryMode:           "async",
			HealthCheckInterval: 15000,
		},
		Ssl: &freezonex_openiiot_api.SSLConfig{
			Enable: false,
		},
		Url: "http://tdengine-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:6041/rest/sql",
	}
}

func emqxreplace2(name string) *freezonex_openiiot_api.BridgeConfig {
	return &freezonex_openiiot_api.BridgeConfig{
		Name:             "water_meter",
		Type:             "webhook",
		Body:             "INSERT INTO demo.${payload.tableName} USING demo.water_meter TAGS('${payload.nodeId}', '${payload.area}') VALUES(now,${payload.water_usage},${payload.vibration_data},'${payload.status}')",
		ConnectTimeout:   "15s",
		Enable:           true,
		EnablePipelining: 100,
		Headers: map[string]string{
			"accept":        "application/json",
			"cache-control": "no-cache",
			"connection":    "keep-alive",
			"content-type":  "application/json",
			"keep-alive":    "timeout=5",
			"Authorization": "Basic cm9vdDp0YW9zZGF0YQ==",
		},
		LocalTopic: "emqx_http/#",
		MaxRetries: 3,
		Method:     "post",
		PoolSize:   4,
		PoolType:   "random",
		ResourceOpts: &freezonex_openiiot_api.ResourceOpts{ // Instantiating and assigning a pointer
			WorkerPoolSize:      1,
			MaxBufferBytes:      104857600,
			InflightWindow:      100,
			QueryMode:           "async",
			HealthCheckInterval: 15000,
		},
		Ssl: &freezonex_openiiot_api.SSLConfig{
			Enable: false,
		},
		Url: "http://tdengine-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:6041/rest/sql",
	}
}

func grafanareplace1(name string) *freezonex_openiiot_api.DataSource {
	return &freezonex_openiiot_api.DataSource{
		Id:              0,
		Name:            "x",
		Type:            "grafana-mqtt-datasource",
		TypeLogoUrl:     "/public/plugins/grafana-mqtt-datasource/img/mqtt.svg",
		Url:             "tcp://emqx-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:1883",
		Access:          "proxy",
		ReadOnly:        false,
		Database:        "",
		User:            "",
		OrgId:           1,
		IsDefault:       false,
		BasicAuth:       true,
		BasicAuthUser:   "admin",
		WithCredentials: false,
		JsonData:        "{\"uri\":\"tcp://emqx-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:1883\",\"username\":\"admin\", \"password\":\"public\"}",
		SecureJsonData:  "null",
		Version:         0,
	}
}

func grafanareplace2(name string) *freezonex_openiiot_api.DataSource {
	return &freezonex_openiiot_api.DataSource{
		Id:              0,
		Name:            "tdengine-datasource",
		Type:            "tdengine-datasource",
		TypeLogoUrl:     "/" + name + "/grafana/public/plugins/tdengine-datasource/img/taosdata_logo.png",
		Url:             "http://tdengine-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:6041",
		Access:          "proxy",
		ReadOnly:        false,
		Database:        "",
		User:            "",
		OrgId:           1,
		IsDefault:       true,
		BasicAuth:       true,
		BasicAuthUser:   "admin",
		WithCredentials: false,
		JsonData:        "{\"uri\":\"http://tdengine-openiiot-" + name + ".openiiot-shimu.svc.cluster.local:6041\",\"username\":\"root\", \"password\":\"taosdata\"}",
		SecureJsonData:  "null",
		Version:         0,
	}
}
