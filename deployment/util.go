package deployment

import (
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
)

var (
	deploymentKind       = reflect.TypeOf(appsv1.Deployment{}).Name()
	deploymentAPIVersion = appsv1.SchemeGroupVersion.String()
	labelKeyRun = "run"
)

func GetMetaData() metav1.TypeMeta{
	return metav1.TypeMeta{
		Kind:       deploymentKind,
		APIVersion: deploymentAPIVersion,
	}
}

func GenerateDeployment(name, namespace, image string) appsv1.Deployment{
	if namespace == ""{
		namespace = "default"
	}

	labels := map[string]string{
		labelKeyRun: name,
	}

	return appsv1.Deployment{
		TypeMeta: GetMetaData(),
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name: name,
							Image: image,
						},
					},
				},
			},
		},
	}
}
