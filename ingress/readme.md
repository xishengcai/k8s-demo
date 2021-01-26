# 使用帮助

## require
- macos or linux
- has kubeconfig in ~/.kube/config

## mkdir certs
```shell script
sh ../hack/generate-cert.sh --CN www.test.com --dir=certs
```

## install nginx-controller and backend
```shell script
kubectl apply -f ./
```

## install http ingress
```shell script
kubectl apply -f only-http
curl -H 'host:www.test.com' nodeIP:{nodePort}/hello
```
