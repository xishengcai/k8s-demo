
kubectl run --namespace=test busybox --rm -ti --image busybox /bin/sh

istioctl kube-inject -f  deployment.yaml |kubectl  apply  -f  -

[inject](https://istio.io/latest/zh/blog/2019/data-plane-setup/)
