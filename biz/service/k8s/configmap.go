package k8s

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) GrafanaConfigMap(namespace string) *corev1.ConfigMap {
	configMapName := "grafana-config"

	return &corev1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "ConfigMap",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      configMapName,
			Namespace: namespace,
		},
		Data: map[string]string{
			"grafana.ini": "[server]\nroot_url = %(protocol)s://%(domain)s:%(http_port)s/" + namespace + "/grafana/",
		},
	}
}

func (a *K8sService) CreateGrafanaConfigMap(ctx context.Context, namespace string) error {
	configMap := a.GrafanaConfigMap(namespace)

	// Create the ConfigMap
	_, err := a.clientset.CoreV1().ConfigMaps(namespace).Create(ctx, configMap, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create ConfigMap: %w", err)
	}

	return nil
}
