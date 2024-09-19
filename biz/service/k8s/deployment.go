package k8s

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment creates a deployment based on the provided deployment name.
func (a *K8sService) CreateDeployment(ctx context.Context, k8sUns K8sUns) error {
	deployment, err := a.getDeploymentSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get deployment spec: %w", err)
	}

	// Create the deployment
	_, err = a.clientset.AppsV1().Deployments(a.GetNamespaceName(k8sUns)).Create(ctx, deployment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create deployment: %w", err)
	}

	return nil
}

// getDeploymentSpec returns a Deployment spec based on the deployment name
func (a *K8sService) getDeploymentSpec(ctx context.Context, k8sUns K8sUns) (*appsv1.Deployment, error) {

	if k8sUns.DeploymentCategory == "app" {
		switch k8sUns.ComponentType {
		case "frontend":
			return a.ApplicationFrontendDeployment(ctx, k8sUns), nil
		case "backend":
		case "db":
		default:
			return nil, fmt.Errorf("deployment type %s not found", k8sUns.ComponentType)
		}
	}

	switch k8sUns.ComponentName {
	case "consolemanager":
		return a.ConsolemanagerDeployment(ctx, k8sUns), nil
	case "emqx":
		return a.EmqxDeployment(ctx, k8sUns), nil
	case "grafana":
		return a.GrafanaDeployment(ctx, k8sUns), nil
	case "mysql":
		return a.MysqlDeployment(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedDeployment(ctx, k8sUns), nil
	case "pma":
		return a.PmaDeployment(ctx, k8sUns), nil
	case "server":
		return a.ServerDeployment(ctx, k8sUns), nil
	case "tdengine":
		return a.TdengineDeployment(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("deployment %s not found", a.GetDeploymentName(k8sUns))
	}
}

// StopDeployment scales the deployment down to zero replicas
func (a *K8sService) StopDeployment(ctx context.Context, k8sUns K8sUns) error {
	scale := int32(0)
	return a.scaleDeployment(ctx, k8sUns, &scale)
}

// StartDeployment scales the deployment up to a specified number of replicas
func (a *K8sService) StartDeployment(ctx context.Context, k8sUns K8sUns) error {
	scale := int32(1)
	return a.scaleDeployment(ctx, k8sUns, &scale)
}

// RestartDeployment restarts the deployment
func (a *K8sService) RestartDeployment(ctx context.Context, k8sUns K8sUns) error {
	return a.clientset.AppsV1().Deployments(a.GetNamespaceName(k8sUns)).Delete(ctx, a.GetDeploymentName(k8sUns), metav1.DeleteOptions{
		PropagationPolicy: func() *metav1.DeletionPropagation {
			fg := metav1.DeletePropagationForeground
			return &fg
		}(),
	})
}

// DeleteDeployment deletes the deployment
func (a *K8sService) DeleteDeployment(ctx context.Context, k8sUns K8sUns) error {
	deploymentName := a.GetDeploymentName(k8sUns)
	namespaceName := a.GetNamespaceName(k8sUns)

	// Delete the Deployment
	err := a.clientset.AppsV1().Deployments(namespaceName).Delete(ctx, deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete deployment %s: %w", deploymentName, err)
	}

	// Wait until the Deployment is fully deleted
	for {
		// Check if the Deployment still exists
		_, err := a.clientset.AppsV1().Deployments(namespaceName).Get(ctx, deploymentName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// Deployment is fully deleted
				break
			}
			return fmt.Errorf("failed to get deployment status: %w", err)
		}

		// Deployment is still terminating, wait and retry
		time.Sleep(2 * time.Second)
	}

	return nil
}

// scaleDeployment is a helper function to scale a deployment
func (a *K8sService) scaleDeployment(ctx context.Context, k8sUns K8sUns, replicas *int32) error {
	namespaceName := a.GetNamespaceName(k8sUns)
	deployment, err := a.clientset.AppsV1().Deployments(namespaceName).Get(ctx, a.GetDeploymentName(k8sUns), metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	deployment.Spec.Replicas = replicas
	_, err = a.clientset.AppsV1().Deployments(namespaceName).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to scale deployment: %w", err)
	}

	return nil
}

// GetAllDeployments returns all deployments across all namespaces.
func (a *K8sService) GetDeploymentInAllNamespaces(ctx context.Context) ([]appsv1.Deployment, error) {
	// Use an empty string as the namespace to list deployments across all namespaces
	deployments, err := a.clientset.AppsV1().Deployments("").List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments across all namespaces: %w", err)
	}

	return deployments.Items, nil
}

// GetDeployment returns all deployments in the given namespace.
func (a *K8sService) GetDeployment(ctx context.Context, k8sUns K8sUns) ([]appsv1.Deployment, error) {
	namespaceName := a.GetNamespaceName(k8sUns)
	deployments, err := a.clientset.AppsV1().Deployments(namespaceName).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments in namespace %s: %w", namespaceName, err)
	}

	return deployments.Items, nil
}

// GetDeploymentByName returns all deployments in the given namespace which their name exact match deploymentName
func (a *K8sService) GetDeploymentByName(ctx context.Context, k8sUns K8sUns) (*appsv1.Deployment, error) {
	namespaceName := a.GetNamespaceName(k8sUns)
	deploymentName := a.GetDeploymentName(k8sUns)
	deployment, err := a.clientset.AppsV1().Deployments(namespaceName).Get(ctx, deploymentName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments in namespace %s: %w", namespaceName, err)
	}

	return deployment, nil
}

// GetDeploymentsByFuzzyName returns all deployments in the given namespace whose names include the given deploymentName.
func (a *K8sService) GetDeploymentsByFuzzyName(ctx context.Context, k8sUns K8sUns) ([]appsv1.Deployment, error) {
	namespaceName := a.GetNamespaceName(k8sUns)
	deploymentName := a.GetDeploymentName(k8sUns)

	// List all deployments in the given namespace
	deployments, err := a.clientset.AppsV1().Deployments(namespaceName).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments in namespace %s: %w", namespaceName, err)
	}

	// Filter deployments whose names contain the given deploymentName (fuzzy match)
	var filteredDeployments []appsv1.Deployment
	for _, deployment := range deployments.Items {
		if strings.Contains(deployment.Name, deploymentName) {
			filteredDeployments = append(filteredDeployments, deployment)
		}
	}

	/*if len(filteredDeployments) == 0 {
		return nil, fmt.Errorf("no deployments found matching the name %s", deploymentName)
	}*/

	return filteredDeployments, nil
}

// GetDeploymentNamesInNamespace returns the names of all deployments in the given namespace.
func (a *K8sService) GetDeploymentNames(ctx context.Context, k8sUns K8sUns) ([]string, error) {
	namespaceName := a.GetNamespaceName(k8sUns)
	deployments, err := a.clientset.AppsV1().Deployments(namespaceName).List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments in namespace %s: %w", namespaceName, err)
	}

	// Extract deployment names
	var deploymentNames []string
	for _, deployment := range deployments.Items {
		deploymentNames = append(deploymentNames, deployment.Name)
	}

	return deploymentNames, nil
}

// GetNextAvailableComponentDeploymentNumber returns the next available number for a specific component, ignoring leading words.
// If no deployments are found for the component, it returns 1.
func (a *K8sService) GetNextAvailableComponentDeploymentNumber(ctx context.Context, k8sUns K8sUns) (int, error) {
	deploymentNames, err := a.GetDeploymentNames(ctx, k8sUns)
	if err != nil {
		return 0, fmt.Errorf("failed to list deployments: %w", err)
	}

	// Dynamically create a regex pattern using the componentName
	pattern := regexp.MustCompile(fmt.Sprintf(`%s(\d+)$`, regexp.QuoteMeta(k8sUns.ComponentName)))

	var numbers []int
	for _, name := range deploymentNames {
		if matches := pattern.FindStringSubmatch(name); matches != nil {
			if num, err := strconv.Atoi(matches[1]); err == nil {
				numbers = append(numbers, num)
			}
		}
	}

	// If no numbers were found, return 1
	if len(numbers) == 0 {
		return 1, nil
	}

	// Sort the numbers
	sort.Ints(numbers)

	// Find the first missing number
	for i, num := range numbers {
		if num != i+1 {
			return i + 1, nil
		}
	}

	// If no number is missing, return the next number
	return len(numbers) + 1, nil
}

// UpdateDeployment updates the deployment to add alias to metadata->labels, spec->template->metadata->labels.
/*func (a *K8sService) UpdateDeployment(ctx context.Context, k8sUns K8sUns) error {
	namespaceName := a.GetNamespaceName(k8sUns)
	deployment, err := a.clientset.AppsV1().Deployments(namespaceName).Get(ctx, a.GetDeploymentName(k8sUns), metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	// Add or update the alias label in metadata->labels
	if deployment.ObjectMeta.Labels == nil {
		deployment.ObjectMeta.Labels = make(map[string]string)
	}
	deployment.ObjectMeta.Labels["alias"] = k8sUns.Alias

	// Add or update the alias label in spec->template->metadata->labels
	if deployment.Spec.Template.ObjectMeta.Labels == nil {
		deployment.Spec.Template.ObjectMeta.Labels = make(map[string]string)
	}
	deployment.Spec.Template.ObjectMeta.Labels["alias"] = k8sUns.Alias

	// Update the deployment with the new labels
	_, err = a.clientset.AppsV1().Deployments(namespaceName).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment: %w", err)
	}

	return nil
}*/

// UpdateDeployment updates the deployment to add alias to metadata->annotations, spec->template->metadata->annotations.
func (a *K8sService) UpdateDeployment(ctx context.Context, k8sUns K8sUns) error {
	namespaceName := a.GetNamespaceName(k8sUns)
	deployment, err := a.clientset.AppsV1().Deployments(namespaceName).Get(ctx, a.GetDeploymentName(k8sUns), metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("failed to get deployment: %w", err)
	}

	// Add or update the alias annotation in metadata->annotations
	if deployment.ObjectMeta.Annotations == nil {
		deployment.ObjectMeta.Annotations = make(map[string]string)
	}
	deployment.ObjectMeta.Annotations["alias"] = k8sUns.Alias

	// Add or update the alias annotation in spec->template->metadata->annotations
	if deployment.Spec.Template.ObjectMeta.Annotations == nil {
		deployment.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	deployment.Spec.Template.ObjectMeta.Annotations["alias"] = k8sUns.Alias

	// Update the deployment with the new annotations
	_, err = a.clientset.AppsV1().Deployments(namespaceName).Update(ctx, deployment, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("failed to update deployment: %w", err)
	}

	return nil
}
