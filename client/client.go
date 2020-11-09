package client

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type Client struct{
	// RESTClient是最基础的客户端RESTClient对HTTP Request进行了封装，实现了RESTful风格的API。
	// ClientSet，DynamicClient，DiscoveryClient客户端都是基于RESTClient实现的。
	RestClient      rest.RESTClient

	// ClientSet 是在RESTClient基础上封装了对Resource和Version的管理方法。
	// 每一个Resource可以理解为一个客户端，而ClientSet则是多个客户端的集合，每一个Resource和Version
	// 都以函数的方式暴露给开发者。ClientSet只能够处理Kubernetes内置资源，他是通过Client-go代码生成器生成的。
	ClientSet      *kubernetes.Clientset

	// DynamicClient与ClientSet最大的不同之处是，ClientSet仅能访问Kubernetes自带的资源（即client集合哪的资源），
	// 而不能直接访问CRD自带的资源。DynamicClient能过处理Kubernetes中的所有资源对象，包括Kubernetes内置资源与
	// CRD自定义资源。
	DynamicClient   dynamic.Interface

	// 发现客户端，用于发现kube-apiserver所支持的资源组、资源版本、资源信息（即Group, Versions,Resources)
	DiscoveryClient *discovery.DiscoveryClient

	RestConfig      *rest.Config
}


