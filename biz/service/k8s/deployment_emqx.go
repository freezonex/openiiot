package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) EmqxDeployment(ctx context.Context, k8sUns K8sUns) *appsv1.Deployment {

	deploymentName := a.GetDeploymentName(k8sUns)

	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: a.GetNamespaceName(k8sUns),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deploymentName,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deploymentName,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentName,
							Image: "openiiot-emqx:1.0.0",
							Env: []corev1.EnvVar{
								{
									Name:  "EMQX_NODE_NAME",
									Value: "emqx@node1.emqx.io",
								},
								{
									Name:  "EMQX_CLUSTER__DISCOVERY_STRATEGY",
									Value: "static",
								},
								{
									Name:  "EMQX_CLUSTER__STATIC__SEEDS",
									Value: "[emqx@node1.emqx.io,emqx@node2.emqx.io]",
								},
							},
							Ports: []corev1.ContainerPort{
								{ContainerPort: 1883},
								{ContainerPort: 8083},
								{ContainerPort: 8084},
								{ContainerPort: 8883},
								{ContainerPort: 18083},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceCPU:    resourceQuantity("10m"),
									corev1.ResourceMemory: resourceQuantity("128Mi"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceCPU:    resourceQuantity("1000m"),
									corev1.ResourceMemory: resourceQuantity("2Gi"),
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "emqx-data",
									MountPath: "/opt/emqx/data",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "emqx-data",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: a.GetPersistentVolumeClaimName(k8sUns),
								},
							},
						},
					},
				},
			},
		},
	}
}
