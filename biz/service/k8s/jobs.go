package k8s

func CreateInstallGrafanaPluginsJob(name string) Job {
	ttl := int32(100) // TTL 为 100 秒

	job := Job{
		APIVersion: "batch/v1",
		Kind:       "Job",
		Metadata: Metadata{
			Name:      "install-grafana-plugins",
			Namespace: name,
		},
		Spec: JobSpec{
			TTLSecondsAfterFinished: &ttl,
			Template: PodTemplateSpec{
				Spec: PodSpec{
					Containers: []Container{
						{
							Name:    "install-plugins",
							Image:   "openiiot_grafana:1.0.0",
							Command: []string{"/bin/sh", "-c", "/opt/grafana/install-plugins.sh"},
							VolumeMounts: []VolumeMount{
								{
									Name:      "plugin-volume",
									MountPath: "/var/lib/grafana",
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []Volume{
						{
							Name: "plugin-volume",
							PersistentVolumeClaim: PersistentVolumeClaim{
								ClaimName: "openiiot-grafana-data-pvc-" + name,
							},
						},
					},
				},
			},
		},
	}

	return job
}
