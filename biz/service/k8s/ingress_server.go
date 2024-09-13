package k8s

import (
	"context"

	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (a *K8sService) ServerIngress(ctx context.Context, k8sUns K8sUns) *networkingv1.Ingress {

	ingressName := a.GetIngressName(k8sUns)

	return &networkingv1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.k8s.io/v1",
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: a.GetNamespaceName(k8sUns),
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
				"nginx.ingress.kubernetes.io/use-regex":      "true",
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: strPtr("nginx"),
			Rules: []networkingv1.IngressRule{
				{
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     a.GetIngressPath(k8sUns),
									PathType: pathTypePtr(networkingv1.PathTypeImplementationSpecific),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: a.GetServicName(k8sUns),
											Port: networkingv1.ServiceBackendPort{
												Number: 8085,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
