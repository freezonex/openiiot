package k8s

import (
	"context"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) GrafanaJob(ctx context.Context, k8sUns K8sUns) *batchv1.Job {
	jobName := a.GetJobName(k8sUns)

	return &batchv1.Job{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "batch/v1",
			Kind:       "Job",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      jobName,
			Namespace: a.GetNamespaceName(k8sUns),
		},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: int32Ptr(100),
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "install-plugins",
							Image:   "openiiot-grafana:1.0.0",
							Command: []string{"/bin/sh", "-c", "/opt/grafana/install-plugins.sh"},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "plugin-volume",
									MountPath: "/var/lib/grafana",
								},
							},
						},
					},
					RestartPolicy: corev1.RestartPolicyNever,
					Volumes: []corev1.Volume{
						{
							Name: "plugin-volume",
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
