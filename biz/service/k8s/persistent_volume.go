package k8s

func NoderedPersistentVolume(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-nodered-volume-" + name,
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
				Path: "/volumes/nfs/openiiot/nodered_data/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func GrafanaPersistentVolume(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-grafana-volume-" + name,
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
				Path: "/volumes/nfs/openiiot/grafana_data/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeData(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-tdengine-volume-data-" + name,
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
				Path: "/volumes/nfs/openiiot/tdengine_data/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func TdenginePersistentVolumeLog(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-tdengine-volume-log-" + name,
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
				Path: "/volumes/nfs/openiiot/tdengine_log/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func EmqxPersistentVolume(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-emqx-volume-" + name,
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
				Path: "/volumes/nfs/openiiot/emqx_data/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}

func MysqlPersistentVolume(name string) *PersistentVolume {
	return &PersistentVolume{
		APIVersion: "v1",
		Kind:       "PersistentVolume",
		Metadata: struct {
			Name string `json:"name"`
		}{
			Name: "openiiot-mysql-volume-" + name,
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
				Path: "/volumes/nfs/openiiot/mysql_data/" + name,
			},
			StorageClassName: "managed-nfs-storage",
		},
	}
}
