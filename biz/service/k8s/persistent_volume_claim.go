package k8s

func NoderedPersistentVolumeClaim(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-nodered-pvc-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-nodered-volume-" + name + "-" + id, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func GrafanaPersistentVolumeClaim(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-grafana-pvc-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-grafana-volume-" + name + "-" + id,
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeClaimData(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-tdengine-pvc-data-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-tdengine-volume-data-" + name + "-" + id,
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeClaimLog(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-tdengine-pvc-log-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-tdengine-volume-log-" + name + "-" + id,
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func EmqxPersistentVolumeClaim(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-emqx-pvc-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-emqx-volume-" + name + "-" + id,
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func MysqlPersistentVolumeClaim(name string, id string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-mysql-pvc-" + name,
		},
		Spec: PersistentVolumeClaimSpec{
			AccessModes: []string{"ReadWriteMany"},
			Resources: struct {
				Requests struct {
					Storage string `json:"storage"`
				} `json:"requests"`
			}{
				Requests: struct {
					Storage string `json:"storage"`
				}{
					Storage: "1Gi",
				},
			},
			VolumeName:       "openiiot-mysql-volume-" + name + "-" + id,
			StorageClassName: "managed-nfs-storage",
		},
	}
}
