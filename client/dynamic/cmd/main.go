package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
)
const (
	OamComponentLabelKey = "oam.runtime.component.id"
	workspace = "366d8f18-b5f7-481e-8522-48fb3b775f14"
	componentID = "393103cd-73ad-481c-b6c2-38a8d6d1aa7b"

)
var (
	nameSpaceResource = schema.GroupVersionResource{Version: "v1", Resource: "namespaces"}
	podGroupVersionResource = schema.GroupVersionResource{Group: "", Version: "v1", Resource: "pods"}
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
			Annotations: map[string]string{
				"a":"b",
			},
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

	nsList, err := dynamicClient.Resource(nameSpaceResource).List(context.Background(),metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	klog.Infof("list :%d", len(nsList.Items))

	respData, err := dynamicClient.Resource(nameSpaceResource).Create(context.Background(),&unstructuredObj,metav1.CreateOptions{})
	if err != nil {
		klog.Fatal(err)
	}
	dynamicClient.Resource(nameSpaceResource).Update(context.Background(), &unstructuredObj, metav1.UpdateOptions{})

	respNs := &corev1.Namespace{}
	// convert unstructured.Unstructured to a Node
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(respData.UnstructuredContent(), respNs); err != nil {
		klog.Fatal(err)
	}

	klog.Infof("namespace: %+v", respNs)
	listOptions := metav1.ListOptions{LabelSelector: fmt.Sprintf("%s=%s", OamComponentLabelKey, componentID)}

	podList, err := ListPods(dynamicClient, listOptions, workspace)
	if err != nil{
		klog.Fatal(err)
	}

	klog.Info(podList)
}

func ListPods(c dynamic.Interface, option metav1.ListOptions, ns string) (*corev1.PodList, error) {
	unstructuredList, err :=c.Resource(podGroupVersionResource).Namespace(ns).List(context.Background(),option)
	if err != nil {
		return nil, err
	}
	podList := &corev1.PodList{}
	if err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructuredList.UnstructuredContent(), podList); err != nil {
		klog.Fatal(err)
	}
	return podList, nil
}

func PatchNs(ns corev1.Namespace, c dynamic.Interface){
	c.Resource(nameSpaceResource).Update(context.Background(), &ns, metav1.UpdateOptions{})

}
