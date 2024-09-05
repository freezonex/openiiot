package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) GrafanaPersistentVolume(ctx context.Context, k8sUns K8sUns) *corev1.PersistentVolume {

	pvName := a.GetPersistentVolumeName(k8sUns)
	pvPath := a.GetPersistentVolumePath(k8sUns)

	return &corev1.PersistentVolume{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "PersistentVolume",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: pvName,
			Labels: map[string]string{
				"type": pvName,
			},
		},
		Spec: corev1.PersistentVolumeSpec{
			Capacity: corev1.ResourceList{
				corev1.ResourceStorage: resourceQuantity("1Gi"),
			},
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteMany,
			},
			PersistentVolumeSource: corev1.PersistentVolumeSource{
				HostPath: &corev1.HostPathVolumeSource{
					Path: pvPath,
				},
			},
		},
	}
}
