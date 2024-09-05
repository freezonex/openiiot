package k8s

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) ServerDeployment(ctx context.Context, k8sUns K8sUns) *appsv1.Deployment {

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
							Name:            deploymentName,
							Image:           "openiiot-server:1.0.0",
							ImagePullPolicy: corev1.PullAlways,
							Env: []corev1.EnvVar{
								{
									Name:  "RUNTIME_IDC_NAME",
									Value: "sg",
								},
								{
									Name:  "SUPOS_SUPOS_ADDRESS",
									Value: "http://47.236.10.165:8080/",
								},
								{
									Name:  "SUPOS_APP_ACCOUNT_ID",
									Value: "c01f69318340143ce7cc9c3870173fa0",
								},
								{
									Name:  "SUPOS_SUPOS_APP_SECRET_KEY",
									Value: "5d6e241c7962badbc630e1aae6d25a28",
								},
							},
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 8085,
								},
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
									Name:      "server-data",
									MountPath: "/volumes",
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "server-data",
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
