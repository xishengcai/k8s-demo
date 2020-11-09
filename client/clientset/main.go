package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
)

func main() {
	config := ctrl.GetConfigOrDie()
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("error")
	}
	deploys, err := c.AppsV1().Deployments("default").List(context.TODO(),v1.ListOptions{})
	if err != nil {
		fmt.Println("error")
	}
	for _, d := range deploys.Items {
		fmt.Println(d.Name)
	}
}
