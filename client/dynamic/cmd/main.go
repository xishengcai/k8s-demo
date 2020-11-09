package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
)


func main(){

	config := ctrl.GetConfigOrDie()

	// create the dynamic client from kubeconfig
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-1",
		},
	}
	// convert the runtime.Object to unstructured.Unstructured
	mapData, err := runtime.DefaultUnstructuredConverter.ToUnstructured(ns)
	if err != nil {
		klog.Fatal(err)
	}
	unstructuredObj := unstructured.Unstructured{
		Object: mapData,
	}

	// create the object using the dynamic client
	nameSpaceResource := schema.GroupVersionResource{Version: "v1", Resource: "namespaces"}
	nsList, err := dynamicClient.Resource(nameSpaceResource).List(context.Background(),metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("list :%d", len(nsList.Items))

	respData, err := dynamicClient.Resource(nameSpaceResource).Create(context.Background(),&unstructuredObj,metav1.CreateOptions{})
	if err != nil {
		klog.Fatal(err)
	}

	respNs := &corev1.Namespace{}
	// convert unstructured.Unstructured to a Node
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(respData.UnstructuredContent(), respNs); err != nil {
		klog.Fatal(err)
	}

	klog.Infof("namespace: %+v", respNs)
}
