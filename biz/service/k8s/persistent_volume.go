package k8s

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreatePersistentVolume creates a persistentVolume based on the provided persistentVolume name.
func (a *K8sService) CreatePersistentVolume(ctx context.Context, k8sUns K8sUns) error {

	runtime_idc_name := os.Getenv("RUNTIME_IDC_NAME")
	if runtime_idc_name != "ci" && runtime_idc_name != "local" {
		if err := a.CreatePersistentVolumePath(ctx, k8sUns); err != nil {
			return err
		}
	}

	persistentVolume, err := a.getPersistentVolumeSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get persistentVolume spec: %w", err)
	}

	// Create the persistentVolume
	_, err = a.clientset.CoreV1().PersistentVolumes().Create(ctx, persistentVolume, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create persistentVolume: %w", err)
	}

	return nil
}

// getPersistentVolumeSpec returns a PersistentVolume spec based on the persistentVolume name
func (a *K8sService) getPersistentVolumeSpec(ctx context.Context, k8sUns K8sUns) (*corev1.PersistentVolume, error) {

	componentName := k8sUns.ComponentName
	if k8sUns.Tag != "" {
		componentName = componentName + "-" + k8sUns.Tag
	}

	switch componentName {
	case "emqx":
		return a.EmqxPersistentVolume(ctx, k8sUns), nil
	case "grafana":
		return a.GrafanaPersistentVolume(ctx, k8sUns), nil
	case "mysql":
		return a.MysqlPersistentVolume(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedPersistentVolume(ctx, k8sUns), nil
	case "tdengine-data":
		return a.TdengineDataPersistentVolume(ctx, k8sUns), nil
	case "tdengine-log":
		return a.TdengineLogPersistentVolume(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("persistentVolume definition for component %s not found", componentName)
	}
}

// delte single pv
func (a *K8sService) DeletePersistentVolume(ctx context.Context, k8sUns K8sUns) error {

	persistentVolumeName := a.GetPersistentVolumeName(k8sUns)

	// Delete the persistentVolume
	err := a.clientset.CoreV1().PersistentVolumes().Delete(ctx, persistentVolumeName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete persistentVolume %s: %w", persistentVolumeName, err)
	}

	// Wait until the PersistentVolume is fully deleted
	for {
		// Check if the PersistentVolume still exists
		_, err := a.clientset.CoreV1().PersistentVolumes().Get(ctx, persistentVolumeName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// PersistentVolume is fully deleted
				break
			}
			return fmt.Errorf("failed to get PersistentVolume status: %w", err)
		}

		// PersistentVolume is still terminating, wait and retry
		time.Sleep(2 * time.Second)
	}

	runtime_idc_name := os.Getenv("RUNTIME_IDC_NAME")
	if runtime_idc_name != "ci" && runtime_idc_name != "local" {
		if err := a.DeletePersistentVolumePath(ctx, k8sUns); err != nil {
			return err
		}
	}

	return nil
}

// delete all pv in 1 tenant
func (a *K8sService) DeleteTenantPersistentVolume(ctx context.Context, tenantName string) error {
	pvNamePrefix := "openiiot-" + tenantName

	// List all PersistentVolumes
	pvList, err := a.clientset.CoreV1().PersistentVolumes().List(ctx, metav1.ListOptions{})
	if err != nil {
		return fmt.Errorf("failed to list PersistentVolumes: %v", err)
	}

	// Iterate through the PVs and delete those that match the prefix
	for _, pv := range pvList.Items {
		if strings.HasPrefix(pv.Name, pvNamePrefix) {
			err := a.clientset.CoreV1().PersistentVolumes().Delete(context.TODO(), pv.Name, metav1.DeleteOptions{})
			if err != nil {
				return fmt.Errorf("failed to delete PersistentVolume %s: %v", pv.Name, err)
			}
			fmt.Printf("Deleted PersistentVolume: %s\n", pv.Name)
		}
	}

	return nil
}
