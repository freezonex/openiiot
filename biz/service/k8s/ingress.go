package k8s

func NoderedIngress(name string) *Ingress {
	return &Ingress{
		APIVersion: "networking.k8s.io/v1beta1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name: name + "nodered",
			Annotations: map[string]string{
				//"nginx.ingress.kubernetes.io/configuration-snippet": "proxy_set_header X-Custom-Header '" + name + "';",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
			},
		},
		Spec: IngressSpec{
			Rules: []Rule{
				{
					HTTP: HTTPIngressRuleValue{
						Paths: []Path{
							{
								Path:     "/" + name + "/nodered/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "nodered-openiiot-" + name,
									ServicePort: 1880,
								},
							},
						},
					},
				},
			},
		},
	}
}

func GrafanaIngress(name string) *Ingress {
	return &Ingress{
		APIVersion: "networking.k8s.io/v1beta1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name: name + "grafana",
			Annotations: map[string]string{
				//"nginx.ingress.kubernetes.io/configuration-snippet": "proxy_set_header X-Custom-Header '" + name + "';",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
			},
		},
		Spec: IngressSpec{
			Rules: []Rule{
				{
					HTTP: HTTPIngressRuleValue{
						Paths: []Path{
							{
								Path:     "/" + name + "/grafana/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "grafana-openiiot-" + name,
									ServicePort: 3000,
								},
							},
						},
					},
				},
			},
		},
	}
}

func WebIngress(name string) *Ingress {
	return &Ingress{
		APIVersion: "networking.k8s.io/v1beta1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name: name + "web",
			Annotations: map[string]string{
				//"nginx.ingress.kubernetes.io/configuration-snippet": "proxy_set_header X-Custom-Header '" + name + "';",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
			},
		},
		Spec: IngressSpec{
			Rules: []Rule{
				{
					HTTP: HTTPIngressRuleValue{
						Paths: []Path{
							{
								Path:     "/" + name + "/web/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "web-openiiot-" + name,
									ServicePort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}

func CombinedIngress(name string) *Ingress {
	return &Ingress{
		APIVersion: "networking.k8s.io/v1beta1",
		Kind:       "Ingress",
		Metadata: Metadata{
			Name: name + "-combined",
			Annotations: map[string]string{
				//"nginx.ingress.kubernetes.io/configuration-snippet": "proxy_set_header X-Custom-Header '" + name + "';",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
			},
		},
		Spec: IngressSpec{
			Rules: []Rule{
				{
					HTTP: HTTPIngressRuleValue{
						Paths: []Path{
							{
								Path:     "/" + name + "/nodered/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "nodered-openiiot-" + name,
									ServicePort: 1880,
								},
							},
							{
								Path:     "/" + name + "/grafana/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "grafana-openiiot-" + name,
									ServicePort: 3000,
								},
							},
							{
								Path:     "/" + name + "/web/(.*)",
								PathType: "Prefix",
								Backend: Backend{
									ServiceName: "web-openiiot-" + name,
									ServicePort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}
