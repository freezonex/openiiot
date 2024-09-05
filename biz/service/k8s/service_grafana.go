package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func (a *K8sService) GrafanaService(ctx context.Context, k8sUns K8sUns) *corev1.Service {

	serviceName := a.GetServicName(k8sUns)

	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      serviceName,
			Namespace: a.GetNamespaceName(k8sUns),
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Name:       "grafana3000",
					Port:       3000,
					TargetPort: intstr.FromInt(3000),
				},
			},
			Selector: map[string]string{
				"app": a.GetDeploymentName(k8sUns),
			},
		},
	}
}
