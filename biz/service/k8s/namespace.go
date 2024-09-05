package k8s

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) CreateNamespace(ctx context.Context, k8sUns K8sUns) error {

	namespaceName := a.GetNamespaceName(k8sUns)

	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}

	// Create the namespace
	_, err := a.clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create namespace: %w", err)
	}

	return nil
}

// DeleteNamespace deletes the specified namespace and all resources within it.
func (a *K8sService) DeleteNamespace(ctx context.Context, k8sUns K8sUns) error {

	namespaceName := a.GetNamespaceName(k8sUns)
	// Delete the namespace
	err := a.clientset.CoreV1().Namespaces().Delete(ctx, namespaceName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete namespace: %w", err)
	}

	// Wait until the namespace is fully deleted
	for {
		// Check if the namespace still exists
		_, err := a.clientset.CoreV1().Namespaces().Get(ctx, namespaceName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// Namespace is fully deleted
				break
			}
			return fmt.Errorf("failed to get namespace status: %w", err)
		}

		// Namespace is still terminating, wait and retry
		time.Sleep(2 * time.Second)
	}

	return nil
}

// GetNamespaces retrieves a list of all namespaces in the cluster.
func (a *K8sService) GetNamespaces(ctx context.Context) ([]v1.Namespace, error) {

	// List all namespaces
	namespaceList, err := a.clientset.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list namespaces: %w", err)
	}

	return namespaceList.Items, nil
}

// NamespaceExists checks if a given namespace exists in the cluster.
func (a *K8sService) NamespaceExists(ctx context.Context, namespaceName string) (bool, error) {

	namespaces, err := a.GetNamespaces(ctx)
	if err != nil {
		return false, err
	}

	for _, ns := range namespaces {
		if ns.Name == namespaceName {
			return true, nil
		}
	}

	return false, nil
}
