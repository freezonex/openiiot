package k8s

func NoderedPersistentVolume(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-nodered-volume-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/nodered_data/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func GrafanaDataPersistentVolume(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-grafana-data-volume-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/grafana_data/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func GrafanaConfigPersistentVolume(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-grafana-config-volume-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/grafana_config/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeData(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-tdengine-volume-data-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/tdengine_data/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeLog(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-tdengine-volume-log-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/tdengine_log/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func EmqxPersistentVolume(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-emqx-volume-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/emqx_data/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func MysqlPersistentVolume(name string, id string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-mysql-volume-" + name + "-" + id,
		},
		Spec: PersistentVolumeSpec{
			Capacity: map[string]string{
				"storage": "5Gi",
			},
			AccessModes: []string{
				"ReadWriteMany",
			},
			PersistentVolumeReclaimPolicy: "Retain",
			HostPath: struct {
				Path string `json:"path"`
			}{
				Path: "/volumes/nfs/openiiot/mysql_data/" + name + "-" + id,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}
