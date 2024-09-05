package k8s

import (
	"context"
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateService creates a service based on the provided service name.
func (a *K8sService) CreateService(ctx context.Context, k8sUns K8sUns) error {

	service, err := a.getServiceSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get service spec: %w", err)
	}

	// Create the service
	_, err = a.clientset.CoreV1().Services(a.GetNamespaceName(k8sUns)).Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	return nil
}

// getServiceSpec returns a Service spec based on the service name
func (a *K8sService) getServiceSpec(ctx context.Context, k8sUns K8sUns) (*corev1.Service, error) {

	switch k8sUns.ComponentName {
	case "consolemanager":
		return a.ConsolemanagerService(ctx, k8sUns), nil
	case "emqx":
		return a.EmqxService(ctx, k8sUns), nil
	case "grafana":
		return a.GrafanaService(ctx, k8sUns), nil
	case "mysql":
		return a.MysqlService(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedService(ctx, k8sUns), nil
	case "pma":
		return a.PmaService(ctx, k8sUns), nil
	case "server":
		return a.ServerService(ctx, k8sUns), nil
	case "tdengine":
		return a.TdengineService(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("service %s not found", a.GetServicName(k8sUns))
	}
}

func (a *K8sService) DeleteService(ctx context.Context, k8sUns K8sUns) error {

	serviceName := a.GetServicName(k8sUns)
	namespaceName := a.GetNamespaceName(k8sUns)

	// Delete the service213005

	err := a.clientset.CoreV1().Services(namespaceName).Delete(ctx, serviceName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete service %s: %w", serviceName, err)
	}

	// Wait until the Service is fully deleted
	for {
		_, err := a.clientset.CoreV1().Services(namespaceName).Get(ctx, serviceName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// Service is fully deleted
				break
			}
			return fmt.Errorf("failed to get service status: %w", err)
		}

		// Service is still terminating, wait and retry
		time.Sleep(2 * time.Second)
	}

	return nil
}
