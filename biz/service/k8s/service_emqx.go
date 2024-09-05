package k8s

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

func (a *K8sService) EmqxService(ctx context.Context, k8sUns K8sUns) *corev1.Service {

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
					Name:       "emqx1883",
					Port:       1883,
					TargetPort: intstr.FromInt(1883),
				},
				{
					Name:       "emqx8083",
					Port:       8083,
					TargetPort: intstr.FromInt(8083),
				},
				{
					Name:       "emqx8084",
					Port:       8084,
					TargetPort: intstr.FromInt(8084),
				},
				{
					Name:       "emqx8883",
					Port:       8883,
					TargetPort: intstr.FromInt(8883),
				},
				{
					Name:       "emqx18083",
					Port:       18083,
					TargetPort: intstr.FromInt(18083),
				},
			},
			Selector: map[string]string{
				"app": a.GetDeploymentName(k8sUns),
			},
		},
	}
}
