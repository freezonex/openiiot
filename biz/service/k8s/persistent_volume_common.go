package k8s

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (a *K8sService) CreatePersistentVolumePath(ctx context.Context, k8sUns K8sUns) error {

	pvPath := a.GetPersistentVolumePath(k8sUns)

	// Check if the directory exists
	_, err := os.Stat(pvPath)
	if os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(pvPath, 0777)
		if err != nil {
			return fmt.Errorf("failed to create directory %s: %w", pvPath, err)
		}
	} else if err != nil {
		// Other error, possibly permission issues
		return fmt.Errorf("failed to check directory %s: %w", pvPath, err)
	}

	// Change the directory owner to nobody:nogroup
	err = os.Chown(pvPath, -1, -1)
	if err != nil {
		return fmt.Errorf("failed to change ownership of directory %s: %w", pvPath, err)
	}

	// Change the directory permissions to 777
	err = os.Chmod(pvPath, 0777)
	if err != nil {
		return fmt.Errorf("failed to change permissions of directory %s: %w", pvPath, err)
	}

	// Copy persistent volume data
	if err := a.CopyPersistentVolumeData(ctx, k8sUns); err != nil {
		return fmt.Errorf("failed to copy persistent volume data: %w", err)
	}

	return nil
}

func (a *K8sService) CopyPersistentVolumeData(ctx context.Context, k8sUns K8sUns) error {

	pvPath := a.GetPersistentVolumePath(k8sUns)

	// Define the source tar.gz file path
	tarFilePath := filepath.Join("/volumes/openiiot/server", fmt.Sprintf("%s.tar.gz", k8sUns.ComponentName))

	// Check if the tar.gz file exists
	if _, err := os.Stat(tarFilePath); os.IsNotExist(err) {
		// File does not exist, return nil
		fmt.Printf("File %s does not exist, skipping extraction.\n", tarFilePath)
		return nil
	}

	// Prepare the tar command
	cmd := exec.CommandContext(ctx, "tar", "-xzvf", tarFilePath, "-C", pvPath)

	// Start the command
	fmt.Printf("File %s exist, extract to %s.\n", tarFilePath, pvPath)
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start extraction: %w", err)
	}

	fmt.Printf("Wait for the command to complete")
	// Wait for the command to complete
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("extraction failed: %w", err)
	}

	fmt.Printf("changeOwnershipAndPermissionsRecursively")
	if err := a.changeOwnershipAndPermissionsRecursively(pvPath); err != nil {
		return fmt.Errorf("Error: %w", err)
	}

	// Log success or handle post-extraction steps here if needed
	fmt.Printf("Successfully extracted %s to %s\n", tarFilePath, pvPath)

	return nil
}

func (a *K8sService) changeOwnershipAndPermissionsRecursively(path string) error {
	// Walk through the directory tree
	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed to access path %s: %w", filePath, err)
		}

		// Change ownership to nobody:nogroup
		if err := os.Chown(filePath, -1, -1); err != nil {
			return fmt.Errorf("failed to change ownership of %s: %w", filePath, err)
		}

		// Change permissions to 777
		if err := os.Chmod(filePath, 0777); err != nil {
			return fmt.Errorf("failed to change permissions of %s: %w", filePath, err)
		}

		return nil
	})

	return err
}

func (a *K8sService) DeletePersistentVolumePath(ctx context.Context, k8sUns K8sUns) error {
	pvPath := a.GetPersistentVolumePath(k8sUns)

	// Check if the directory exists
	_, err := os.Stat(pvPath)
	if os.IsNotExist(err) {
		// Directory does not exist, return nil since no action is needed
		return nil
	} else if err != nil {
		// Other error, possibly permission issues
		return fmt.Errorf("failed to check directory %s: %w", pvPath, err)
	}

	// Remove the directory and all its contents
	err = os.RemoveAll(pvPath)
	if err != nil {
		return fmt.Errorf("failed to delete directory %s: %w", pvPath, err)
	}

	return nil
}

func (a *K8sService) DeleteTenantPersistentVolumePath(ctx context.Context, tenantName string) error {
	pvPath := a.GetTenantPersistentVolumePath(tenantName)

	// Check if the directory exists
	_, err := os.Stat(pvPath)
	if os.IsNotExist(err) {
		// Directory does not exist, return nil since no action is needed
		return nil
	} else if err != nil {
		// Other error, possibly permission issues
		return fmt.Errorf("failed to check directory %s: %w", pvPath, err)
	}

	// Remove the directory and all its contents
	err = os.RemoveAll(pvPath)
	if err != nil {
		return fmt.Errorf("failed to delete directory %s: %w", pvPath, err)
	}

	return nil
}
