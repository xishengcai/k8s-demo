package common

import (
	"encoding/json"
	"k8s.io/klog"
)

func PrettifyJson(i interface{}, indent bool) string {
	var str []byte
	var err error
	if indent {
		str, err = json.MarshalIndent(i, "", "    ")
		if err != nil {
			klog.Fatal(err)
		}
	} else {
		str, err = json.Marshal(i)
		if err != nil {
			klog.Fatal(err)
		}
	}

	return string(str)
}

func PrintData(data interface{}, err error) {
	if err != nil {
		klog.Fatal(err)
	}
	klog.Info(PrettifyJson(data, true))
}
