package helloworld

import (
	helloworldv1alpha1 "github.com/BuddhiWathsala/helloworld-k8s-operator/pkg/apis/helloworld/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// newHelloWorldDeployment returns a busybox pod with the same name/namespace as the cr
func newDeployment(cr *helloworldv1alpha1.HelloWorld) *appsv1.Deployment {
	labels := getLabels(cr)
	return &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: appsv1.SchemeGroupVersion.String(),
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &cr.Spec.Size,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:    "helloworld",
							Image:   "docker.io/buddhiwathsala/helloworld-goapp:v0.1.0",
							Command: []string{"go", "run", "hello-world.go"},
						},
					},
				},
			},
		},
	}
}

// newService returns a new K8s service
func newService(cr *helloworldv1alpha1.HelloWorld) *corev1.Service {
	labels := getLabels(cr)
	return &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			APIVersion: corev1.SchemeGroupVersion.String(),
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
		},
		Spec: corev1.ServiceSpec{
			Selector: labels,
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Port: 8080,
					Name: "port8080",
				},
			},
			Type: "ClusterIP",
		},
	}
}

// getLabels return the labels that maps K8s deployment and service
func getLabels(cr *helloworldv1alpha1.HelloWorld) map[string]string {
	labels := map[string]string{
		"app": cr.Name,
	}
	return labels
}
