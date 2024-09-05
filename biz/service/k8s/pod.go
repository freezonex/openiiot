package k8s

import (
	"context"
	"fmt"
	"strings"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) GetRuntimePods(ctx context.Context, k8sUns K8sUns, includeJobPods bool) ([]v1.Pod, error) {

	deploymentName := a.GetDeploymentName(k8sUns)
	namespaceName := a.GetNamespaceName(k8sUns)

	// Define a list options variable
	var listOptions metav1.ListOptions

	// List all Pods in the specified namespace (with or without the label selector)
	pods, err := a.clientset.CoreV1().Pods(namespaceName).List(ctx, listOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to list pods: %w", err)
	}

	var matchedPods []v1.Pod
	// Filter pods based on the deployment name and the includeJobPods flag
	for _, pod := range pods.Items {
		// Check if the pod is owned by a Job
		isJobPod := false
		for _, ownerRef := range pod.OwnerReferences {
			if ownerRef.Kind == "Job" {
				isJobPod = true
				break
			}
		}

		// Include the pod based on the includeJobPods flag and deployment name
		if (includeJobPods || !isJobPod) && (deploymentName == "" || strings.HasPrefix(pod.Name, deploymentName)) {
			matchedPods = append(matchedPods, pod)
		}
	}

	return matchedPods, nil
}
