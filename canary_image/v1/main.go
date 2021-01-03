package main

import (
	"net/http"
	"os"
)

func SayHello(w http.ResponseWriter, req *http.Request) {
	version := os.Getenv("VERSION")
	if version == ""{
		version = "v1"
	}
	w.Write([]byte(version))
}

func main() {
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":80", nil)
}
