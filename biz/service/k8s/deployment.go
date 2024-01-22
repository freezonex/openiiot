package k8s

import "strings"

func NewDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": name,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "registry:5000/dt/javaw-dname:1.0.0",
							Name:  name,
						},
					},
				},
			},
		},
	}
}

func WebDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      "web" + name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": "web-" + name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": "web-" + name,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_web:1.0.0",
							Name:  "web" + name,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
						},
					},
				},
			},
		},
	}
}

func NginxDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      "nginx-" + name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": "nginx-" + name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": "nginx-" + name,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "nginxw:1.0.0",
							Name:  "nginx-" + name,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
						},
					},
				},
			},
		},
	}
}

func NoderedDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      "nodered-" + name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": "nodered-" + name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": "nodered-" + name,
					},
				},
				Spec: PodSpec{
					InitContainers: []Container{
						{
							Name:  "init-chmod",
							Image: "busybox",
							Command: []string{
								"chmod", "-R", "777", "/data",
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "nodered-data",
									MountPath: "/data",
								},
							},
						},
					},
					Containers: []Container{
						{
							Image: "openiiot_nodered:1.0.0",
							Name:  "nodered" + name,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "nodered-data",
									MountPath: "/data",
								},
							},
						},
					},
					//未创建PersistentVolumeClaim会报错
					Volumes: []Volume{
						{
							Name: "nodered-data",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-nodered-pvc-" + name,
							},
						},
					},
				},
			},
		},
	}
}

func ServerDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      "server-" + name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": "server-" + name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": "server-" + name,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_server:1.0.0",
							Name:  "server" + name,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
						},
					},
				},
			},
		},
	}
}

func GrafanaDeployment(name string) *Deployment {
	grafanaName := "grafana-" + name // 组合名字
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      grafanaName,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": grafanaName,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": grafanaName,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_grafana:1.0.0",
							Name:  grafanaName,
							Env: []EnvVar{
								{
									Name:  "GF_SERVER_ROOT_URL", // 设置环境变量
									Value: "%(protocol)s://%(domain)s:%(http_port)s/" + strings.TrimPrefix(name, "openiiot-") + "/grafana/",
								},
							},
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "grafana-data",
									MountPath: "/data",
								},
								//{
								//	Name:      "grafana-config",
								//	MountPath: "/etc/grafana",
								//},
							},
						},
					},
					Volumes: []Volume{
						{
							Name: "grafana-data",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-grafana-data-pvc-" + name,
							},
						},
						//{
						//	Name: "grafana-config",
						//	PersistentVolumeClaim: PersistentVolumeClaim{
						//		ClaimName: "openiiot-grafana-config-pvc-" + name,
						//	},
						//},
					},
				},
			},
		},
	}
}

func TdenggineDeployment(name string) *Deployment {
	tdengineName := "tdengine-" + name // 组合名字
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      tdengineName,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": tdengineName,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": tdengineName,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_tdengine:1.0.0",
							Name:  tdengineName,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "tdengine-data",
									MountPath: "/data",
								},
							},
						},
					},
					//未创建PersistentVolumeClaim会报错
					Volumes: []Volume{
						{
							Name: "tdengine-data",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-tdengine-pvc-data-" + name,
							},
						},
						{
							Name: "tdengine-log",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-tdengine-pvc-log-" + name,
							},
						},
					},
				},
			},
		},
	}
}

func EmqxDeployment(name string) *Deployment {
	emqxName := "emqx-" + name // 组合名字
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      emqxName,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": emqxName,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": emqxName,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_emqx:1.0.0",
							Name:  emqxName,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "emqx-data",
									MountPath: "/data",
								},
							},
						},
					},
					//未创建PersistentVolumeClaim会报错
					Volumes: []Volume{
						{
							Name: "emqx-data",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-emqx-pvc-" + name,
							},
						},
					},
				},
			},
		},
	}
}

func MysqlDeployment(name string) *Deployment {
	mysqlName := "mysql-" + name // 组合名字
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      mysqlName,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": mysqlName,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": mysqlName,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_mysql:1.0.0",
							Name:  mysqlName,
							Env: []EnvVar{ // 添加环境变量
								{
									Name:  "MYSQL_ROOT_PASSWORD",
									Value: "root1234",
								},
								{
									Name:  "MYSQL_DATABASE",
									Value: "openiiot",
								},
								{
									Name:  "MYSQL_USER",
									Value: "openiiot",
								},
								{
									Name:  "MYSQL_PASSWORD",
									Value: "root1234",
								},
							},
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
							VolumeMounts: []VolumeMount{
								{
									Name:      "mysql-data",
									MountPath: "/data",
								},
							},
						},
					},
					//未创建PersistentVolumeClaim会报错
					Volumes: []Volume{
						{
							Name: "mysql-data",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-mysql-pvc-" + name,
							},
						},
					},
				},
			},
		},
	}
}

func ConsolemanagerDeployment(name string) *Deployment {
	return &Deployment{
		APIVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: Metadata{
			Name:      "consolemanager-" + name,
			Namespace: name,
		},
		Spec: Spec{
			Replicas: 1,
			Selector: Selector{
				MatchLabels: map[string]string{
					"app": "consolemanager-" + name,
				},
			},
			Template: PodTemplate{
				Metadata: Metadata{
					Labels: map[string]string{
						"app": "consolemanager-" + name,
					},
				},
				Spec: PodSpec{
					Containers: []Container{
						{
							Image: "openiiot_consolemanager:1.0.0",
							Name:  "consolemanager-" + name,
							Resources: ResourceRequirements{
								Requests: map[string]string{
									"cpu":    "10m",
									"memory": "128Mi",
								},
								Limits: map[string]string{
									"cpu":    "1000m",
									"memory": "2Gi",
								},
							},
						},
					},
				},
			},
		},
	}
}
