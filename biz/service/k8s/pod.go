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

// UpdatePod updates the pods related to the deployment to add alias to metadata->labels
func (a *K8sService) UpdatePod(ctx context.Context, k8sUns K8sUns) error {
	namespaceName := a.GetNamespaceName(k8sUns)

	// Set the ListOptions to filter pods by the deployment's label, e.g., app: <deployment-name>
	listOptions := metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app=%s", a.GetDeploymentName(k8sUns)),
	}

	// Get the list of pods that belong to the deployment by using the label selector
	pods, err := a.clientset.CoreV1().Pods(namespaceName).List(ctx, listOptions)
	if err != nil {
		return fmt.Errorf("failed to list pods: %w", err)
	}

	// Iterate over the list of pods and update the alias label
	for _, pod := range pods.Items {
		// Check if the pod's metadata labels map is nil and initialize if necessary
		if pod.ObjectMeta.Labels == nil {
			pod.ObjectMeta.Labels = make(map[string]string)
		}

		// Add or update the alias label
		pod.ObjectMeta.Labels["alias"] = k8sUns.Alias

		// Update the pod with the new labels
		_, err = a.clientset.CoreV1().Pods(namespaceName).Update(ctx, &pod, metav1.UpdateOptions{})
		if err != nil {
			return fmt.Errorf("failed to update pod %s: %w", pod.Name, err)
		}
	}

	return nil
}
