package k8s

import (
	"context"
	"fmt"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateJob creates a job based on the provided job name.
func (a *K8sService) CreateJob(ctx context.Context, k8sUns K8sUns) error {

	job, err := a.getJobSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get job spec: %w", err)
	}

	// Create the job
	_, err = a.clientset.BatchV1().Jobs(a.GetNamespaceName(k8sUns)).Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create job: %w", err)
	}

	return nil
}

// getJobSpec returns a Job spec based on the job name
func (a *K8sService) getJobSpec(ctx context.Context, k8sUns K8sUns) (*batchv1.Job, error) {

	switch k8sUns.ComponentName {
	case "grafana":
		return a.GrafanaJob(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedJob(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("job %s not found", a.GetJobName(k8sUns))
	}
}

// DeleteJob deletes the job
func (a *K8sService) DeleteJob(ctx context.Context, k8sUns K8sUns) error {

	return a.clientset.BatchV1().Jobs(a.GetNamespaceName(k8sUns)).Delete(ctx, a.GetJobName(k8sUns), metav1.DeleteOptions{})
}
