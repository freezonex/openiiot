package k8s

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreatePersistentVolumeClaim creates a persistentVolumeClaim based on the provided persistentVolumeClaim name.
func (a *K8sService) CreatePersistentVolumeClaim(ctx context.Context, k8sUns K8sUns) error {

	persistentVolumeClaim, err := a.getPersistentVolumeClaimSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get persistentVolumeClaim spec: %w", err)
	}

	// Create the persistentVolumeClaim
	_, err = a.clientset.CoreV1().PersistentVolumeClaims(a.GetNamespaceName(k8sUns)).Create(ctx, persistentVolumeClaim, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create persistentVolumeClaim: %w", err)
	}

	return nil
}

// getPersistentVolumeClaimSpec returns a PersistentVolumeClaim spec based on the persistentVolumeClaim name
func (a *K8sService) getPersistentVolumeClaimSpec(ctx context.Context, k8sUns K8sUns) (*corev1.PersistentVolumeClaim, error) {

	componentName := k8sUns.ComponentName
	if k8sUns.Tag != "" {
		componentName = componentName + "-" + k8sUns.Tag
	}

	switch componentName {
	case "emqx":
		return a.EmqxPersistentVolumeClaim(ctx, k8sUns), nil
	case "grafana":
		return a.GrafanaPersistentVolumeClaim(ctx, k8sUns), nil
	case "mysql":
		return a.MysqlPersistentVolumeClaim(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedPersistentVolumeClaim(ctx, k8sUns), nil
	case "tdengine-data":
		return a.TdengineDataPersistentVolumeClaim(ctx, k8sUns), nil
	case "tdengine-log":
		return a.TdengineLogPersistentVolumeClaim(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("persistentVolumeClaim %s not found", componentName)
	}
}

func (a *K8sService) DeletePersistentVolumeClaim(ctx context.Context, k8sUns K8sUns) error {

	persistentVolumeClaimName := a.GetPersistentVolumeClaimName(k8sUns)
	namespaceName := a.GetNamespaceName(k8sUns)

	// Delete the persistentVolumeClaim
	err := a.clientset.CoreV1().PersistentVolumeClaims(namespaceName).Delete(ctx, persistentVolumeClaimName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete persistentVolumeClaim %s: %w", persistentVolumeClaimName, err)
	}

	// Wait until the PersistentVolumeClaim is fully deleted
	for {
		// Check if the PersistentVolumeClaim still exists
		_, err := a.clientset.CoreV1().PersistentVolumeClaims(namespaceName).Get(ctx, persistentVolumeClaimName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// PersistentVolumeClaim is fully deleted
				break
			}
			return fmt.Errorf("failed to get PersistentVolumeClaim status: %w", err)
		}

		// PersistentVolumeClaim is still terminating, wait and retry
		time.Sleep(2 * time.Second)
	}

	return nil
}
