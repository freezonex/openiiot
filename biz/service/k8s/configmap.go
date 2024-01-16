package k8s

func GrafanaConfigmap(name string) *ConfigMap {
	return &ConfigMap{
		APIVersion: "v1",
		Kind:       "ConfigMap",
		Metadata: Metadata{
			Name: "grafana-config",
		},
		Data: map[string]string{
			"grafana.ini": "[server]root_url = %(protocol)s://%(domain)s:%(http_port)s/" + name + "/grafana/",
		},
	}
}
