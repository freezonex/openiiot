package k8s

func NewService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       80,
					Protocol:   "TCP",
					TargetPort: 8080,
				},
			},
			Selector: map[string]string{
				"app": name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func WebService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "web-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       80,
					Protocol:   "TCP",
					TargetPort: 80,
				},
			},
			Selector: map[string]string{
				"app": "web" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func NoderedService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "nodered-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       1880,
					Protocol:   "TCP",
					TargetPort: 1880,
				},
			},
			Selector: map[string]string{
				"app": "nodered" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func ServerService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "server-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       8085,
					Protocol:   "TCP",
					TargetPort: 8085,
				},
			},
			Selector: map[string]string{
				"app": "server" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func GrafanaService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "grafana-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       3000,
					Protocol:   "TCP",
					TargetPort: 3000,
				},
			},
			Selector: map[string]string{
				"app": "grafana" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func TdengineService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "tdengine-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       6030,
					Protocol:   "TCP",
					TargetPort: 6030,
				},
			},
			Selector: map[string]string{
				"app": "tdengine" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func EmqxService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "emqx-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       1883,
					Protocol:   "TCP",
					TargetPort: 1883,
				},
			},
			Selector: map[string]string{
				"app": "emqx" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func MysqlService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "mysql-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       3306,
					Protocol:   "TCP",
					TargetPort: 3306,
				},
			},
			Selector: map[string]string{
				"app": "mysql" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}

func ConsolemanagerService(name string, nodePort int) *Service {
	return &Service{
		APIVersion: "v1",
		Kind:       "Service",
		Metadata: Metadata{
			Name: "consolemanager-" + name,
		},
		Spec: ServiceSpec{
			Ports: []Port{
				{
					NodePort:   nodePort,
					Port:       81,
					Protocol:   "TCP",
					TargetPort: 81,
				},
			},
			Selector: map[string]string{
				"app": "consolemanager" + name,
			},
			//Type: "ClusterIP",
			Type: "NodePort",
		},
	}
}
