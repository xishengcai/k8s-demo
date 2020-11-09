package main

import (
	"context"
	oamcore "github.com/xishengcai/oam/apis/core"
	oam "github.com/xishengcai/oam/apis/core/v1alpha2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var scheme = runtime.NewScheme()

func init() {
	_ = clientgoscheme.AddToScheme(scheme)
	_ = oamcore.AddToScheme(scheme)
}

func main() {
	c := getInterfaceClient()
	ctx := context.Background()
	comps := &oam.ComponentList{}

	err := c.List(ctx, comps, &client.ListOptions{})
	if err != nil{
		klog.Fatal(err)
	}
	klog.Info(oam.SchemeGroupVersion.String())
	for _, comp := range comps.Items{
		klog.Infof("%+v", comp.Name)
	}

	component := generateOamComponent()

	applyOpts := []client.PatchOption{
		client.ForceOwnership,
		client.FieldOwner("launcher"),
		&client.PatchOptions{FieldManager: "apply"},
	}

	err = c.Patch(context.Background(), component, client.Apply, applyOpts...)
	if err != nil {
		klog.Fatal(err)
	}
	return
}

func getInterfaceClient() client.Client{
	config := ctrl.GetConfigOrDie()
	c, err := client.New(config, client.Options{Scheme: scheme})
	if err != nil {
		klog.Fatal(err)
	}

	return c
}

func generateOamComponent() *oam.Component{
	component := &oam.Component{
		TypeMeta: metav1.TypeMeta{
			Kind:       oam.ComponentKind,
			APIVersion: oam.SchemeGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "oam-sts",
			Namespace: "default",
		},
	}
	return component
}
