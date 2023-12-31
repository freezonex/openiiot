package tenant

func NoderedPersistentVolumeClaim(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-nodered-pvc",
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
			VolumeName:       "openiiot-nodered-volume-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func GrafanaPersistentVolumeClaim(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-grafana-pvc",
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
			VolumeName:       "openiiot-grafana-volume-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeClaimData(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-tdengine-pvc-data",
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
			VolumeName:       "openiiot-tdengine-volume-data-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeClaimLog(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-tdengine-pvc-log",
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
			VolumeName:       "openiiot-tdengine-volume-log-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func EmqxPersistentVolumeClaim(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-emqx-pvc",
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
			VolumeName:       "openiiot-emqx-volume-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func MysqlPersistentVolumeClaim(name string) *PersistentVolumeClaim {
	return &PersistentVolumeClaim{
		APIVersion: "v1",
		Kind:       "PersistentVolumeClaim",
		Metadata: Metadata{
			Name: "openiiot-mysql-pvc",
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
			VolumeName:       "openiiot-mysql-volume-" + name, // 如果不需要指定，可以将这一行删除
			StorageClassName: "managed-nfs-storage",
		},
	}
}
