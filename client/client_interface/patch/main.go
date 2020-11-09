package main

import (
	"context"
	"k8s-demo/deployment"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

/*
use interface client patch one already exist deployment
bug can't apply not exist deployment
 */
func main(){
	config := ctrl.GetConfigOrDie()

	c, err := client.New(config, client.Options{})
	if err != nil {
		klog.Fatal(err)
	}

	dep := &appsv1.Deployment{}
	key := client.ObjectKey{Namespace: "default", Name: "nginx"}
	err = c.Get(context.Background(), key, dep)
	if err != nil {
		klog.Fatal(err)
	}

	dep.SetLabels(map[string]string{"a":"B"})
	dep.TypeMeta = deployment.GetMetaData()
	dep.ObjectMeta.ManagedFields = nil

	applyOpts := []client.PatchOption{client.ForceOwnership, client.FieldOwner(dep.GetUID())}
	err = c.Patch(context.TODO(), dep, client.Apply, applyOpts...)
	if err != nil {
		klog.Fatal(err)
	}
}
