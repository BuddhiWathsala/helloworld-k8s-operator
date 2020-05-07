package helloworld

import (
	helloworldv1alpha1 "github.com/BuddhiWathsala/helloworld-k8s-operator/pkg/apis/helloworld/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/util/intstr"
)

// IntOrString integer or string
type IntOrString struct {
	Type   Type   `protobuf:"varint,1,opt,name=type,casttype=Type"`
	IntVal int32  `protobuf:"varint,2,opt,name=intVal"`
	StrVal string `protobuf:"bytes,3,opt,name=strVal"`
}

// Type represents the stored type of IntOrString.
type Type int

// Int - Type
const (
	Int intstr.Type = iota
	String
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
							Name:            "helloworld",
							Image:           "docker.io/buddhiwathsala/helloworld-goapp:v0.4.0",
							Command:         []string{"go", "run", "hello-world.go"},
							ImagePullPolicy: corev1.PullAlways,
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

// newIngress returns an NGINX ingress
func newNginxIngress(cr *helloworldv1alpha1.HelloWorld) *extensionsv1beta1.Ingress {
	ingressPaths := []extensionsv1beta1.HTTPIngressPath{
		extensionsv1beta1.HTTPIngressPath{
			Path: "/" + cr.Name + "(/|$)(.*)",
			Backend: extensionsv1beta1.IngressBackend{
				ServiceName: cr.Name,
				ServicePort: intstr.IntOrString{
					Type:   Int,
					IntVal: 8080,
				},
			},
		},
	}
	ingressSpec := extensionsv1beta1.IngressSpec{
		Rules: []extensionsv1beta1.IngressRule{
			{
				Host: "helloworld",
				IngressRuleValue: extensionsv1beta1.IngressRuleValue{
					HTTP: &extensionsv1beta1.HTTPIngressRuleValue{
						Paths: ingressPaths,
					},
				},
			},
		},
	}
	ingress := &extensionsv1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			APIVersion: extensionsv1beta1.SchemeGroupVersion.String(),
			Kind:       "Ingress",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cr.Name,
			Namespace: cr.Namespace,
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":                "nginx",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$2",
			},
		},
		Spec: ingressSpec,
	}
	return ingress
}

// getLabels return the labels that maps K8s deployment and service
func getLabels(cr *helloworldv1alpha1.HelloWorld) map[string]string {
	labels := map[string]string{
		"app": cr.Name,
	}
	return labels
}
