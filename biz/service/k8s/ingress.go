package k8s

import (
	"context"
	"fmt"
	"time"

	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateIngress creates a ingress based on the provided ingress name.
func (a *K8sService) CreateIngress(ctx context.Context, k8sUns K8sUns) error {

	ingress, err := a.getIngressSpec(ctx, k8sUns)
	if err != nil {
		return fmt.Errorf("failed to get ingress spec: %w", err)
	}

	// Create the ingress
	_, err = a.clientset.NetworkingV1().Ingresses(a.GetNamespaceName(k8sUns)).Create(ctx, ingress, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create ingress: %w", err)
	}

	return nil
}

// getIngressSpec returns a Ingress spec based on the ingress name
func (a *K8sService) getIngressSpec(ctx context.Context, k8sUns K8sUns) (*networkingv1.Ingress, error) {

	switch k8sUns.ComponentName {
	case "emqx":
		return a.EmqxIngress(ctx, k8sUns), nil
	case "grafana":
		return a.GrafanaIngress(ctx, k8sUns), nil
	case "nodered":
		return a.NoderedIngress(ctx, k8sUns), nil
	default:
		return nil, fmt.Errorf("ingress %s not found", a.GetIngressName(k8sUns))
	}
}

func (a *K8sService) DeleteIngress(ctx context.Context, k8sUns K8sUns) error {
	ingressName := a.GetIngressName(k8sUns)
	namespaceName := a.GetNamespaceName(k8sUns)

	err := a.clientset.NetworkingV1().Ingresses(namespaceName).Delete(ctx, ingressName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete ingress %s: %w", ingressName, err)
	}

	// Wait until the Ingress is fully deleted
	for {
		_, err := a.clientset.NetworkingV1().Ingresses(namespaceName).Get(ctx, ingressName, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				// Ingress is fully deleted
				break
			}
			return fmt.Errorf("failed to get ingress status: %w", err)
		}

		// Ingress is still present, wait and retry
		time.Sleep(2 * time.Second)
	}

	return nil
}
