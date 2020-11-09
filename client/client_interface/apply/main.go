package main

import (
	"context"
	"fmt"
	"k8s-demo/common"
	"k8s-demo/deployment"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

/*
patch deployment
if not found, new one;
else, update.
*/

var (
	workloadName = "apply-test"
	namespace    = "default"
	imageName    = "nginx"
)

func main() {
	ctx := context.Background()
	config := ctrl.GetConfigOrDie()
	begin := time.Now()
	c, err := client.New(config, client.Options{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info("build client cost time: ", time.Since(begin))
	dep := deployment.GenerateDeployment(workloadName, namespace, imageName)

	// 1.delete deployment
	if err := c.Delete(ctx, &dep, &client.DeleteOptions{}); err != nil {
		klog.Fatal(err)
	}

	// patch 选项
	applyOpts := []client.PatchOption{
		client.ForceOwnership,
		client.FieldOwner(dep.GetUID()),
		&client.PatchOptions{FieldManager: "apply"},
	}

	for i:=0;i<3; i++ {
		// 2. 原始数据重复apply
		time.Sleep(1 *time.Second)
		dep.SetAnnotations(map[string]string{fmt.Sprintf("time-%d",i): time.Now().String()})
		dep.ObjectMeta.ManagedFields = nil
		dep.ObjectMeta.ResourceVersion = ""
		err = c.Patch(ctx, &dep, client.Apply, applyOpts...)
		if err != nil {
			klog.Fatal(err)
		}
	}

	// 3.打印当前deployment
	depGet := appsv1.Deployment{}
	key := client.ObjectKey{Namespace: namespace, Name: workloadName}
	if err := c.Get(ctx, key, &depGet); err != nil {
		klog.Fatal(err)
	}
	common.PrintData(depGet, nil)

}
