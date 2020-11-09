package main

import(
	"k8s.io/client-go/tools/clientcmd"
)

func main(){
	config, err := clientcmd.BuildConfigFromFlags("", "~/.kube/config")
	if err != nil{
		panic(err)
	}
}
