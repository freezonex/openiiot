package k8s

import (
	"context"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WaitForJobCompletion waits for the specified Job to complete
func (a *K8sService) WaitForJobCompletion(ctx context.Context, k8sUns K8sUns) error {

	namespaceName := a.GetNamespaceName(k8sUns)
	jobName := a.GetJobName(k8sUns)

	for {
		job, err := a.clientset.BatchV1().Jobs(namespaceName).Get(ctx, jobName, metav1.GetOptions{})
		if err != nil {
			return fmt.Errorf("failed to get Job status: %w", err)
		}

		if job.Status.Succeeded == *job.Spec.Completions {
			fmt.Printf("Job %s completed successfully.\n", jobName)
			break
		}

		fmt.Printf("Waiting for Job %s to complete...\n", jobName)
		time.Sleep(10 * time.Second)
	}

	return nil
}
