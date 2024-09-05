package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) EmqxPersistentVolumeClaim(ctx context.Context, k8sUns K8sUns) *corev1.PersistentVolumeClaim {

	return &corev1.PersistentVolumeClaim{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "PersistentVolumeClaim",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      a.GetPersistentVolumeClaimName(k8sUns),
			Namespace: a.GetNamespaceName(k8sUns),
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteMany,
			},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("1Gi"),
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"type": a.GetPersistentVolumeName(k8sUns),
				},
			},
		},
	}
}
