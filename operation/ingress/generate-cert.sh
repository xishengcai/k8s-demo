#!/bin/bash

usage() {
cat <<EOF

usage:
    --CN common name: localhost
    --IP 1.1.1.1
    --rsa 2048
    -h print help
    example: sh generate-cert.sh --CN localhost

EOF
}

set -e

if [ "$1" = "-h" ]; then
  usage
  exit 0
fi

while [ $# -gt 0 ]
do
    key="$1"
    case $key in
        --IP)
            export IP=$2
            shift
        ;;
        --CN)
            export CN=$2
            shift
        ;;
        --rsa)
            export rsa=$2
            shift
        ;;
        *)
            echo "unknown option [$key]"
            usage
            exit 1
        ;;
    esac
    shift
done

if [ -z "$CN" ]; then
  CN='localhost'
fi

rsa=2048
#if [ -z "$rsa"]; then
#  rsa=2048
#fi

genServerCrt(){
  # 生成根秘钥及证书
  openssl req -x509 -sha256 -newkey rsa:${rsa} -keyout ca.key -out ca.crt -days 3560 -nodes -subj '/CN=Launcher LStack Authority'

  # 生成服务器密钥，证书并使用CA证书签名
  openssl genrsa -out server.key ${rsa}
  openssl req -new -key server.key -subj "/CN=${CN}" -out server.csr

  # ip domain
  if [ -n "$IP" ]; then
    echo "ip: $IP"
    echo subjectAltName = IP:${IP} > extfile.cnf
    openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile extfile.cnf -out server.crt -days 3650
  else
    openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650
  fi
}

# 生成客户端密钥，证书并使用CA证书签名
genClientCert(){
  openssl req -new -newkey rsa:${rsa} -keyout client.key -out client.csr -nodes -subj "/CN=${CN}"
  openssl x509 -req -sha256 -days 3650 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 02 -out client.crt
}

# 生成p12
creteP12(){
  openssl pkcs12 -export -clcerts -inkey client.key -in client.crt -out client.p12
}

createCert() {
  kubectl delete secret tls-secret -n test

  kubectl -n test \
  create secret generic tls-secret \
--from-file=tls.crt=server.crt \
--from-file=tls.key=server.key \
--from-file=ca.crt=ca.crt
}

CreateTls() {
  kubectl delete secret tls-secret -n test
  kubectl -n test \
  create secret tls tls-secret \
--cert=server.crt \
--key=server.key \
}

createCaSecret(){
  kubectl create secret generic ca --from-file=ca.crt=ca.crt -n test
}



testCurl() {
  curl -v https://xxx:6443/apis/launchercontroller.k8s.io/v1/rsoverviews/lau-crd-resource-overview?timeout=5s \
--cert /etc/kubernetes/pki/apiserver-kubelet-client.crt  \
--key /etc/kubernetes/pki/apiserver-kubelet-client.key \
--cacert /etc/kubernetes/pki/ca.crt

curl https://xx \
--cert ./cert/client.crt \
--key ./cert/client.key \
--cacert ./cert/ca.crt
}

clean(){
  rm -rf ./extfile.cnf
  rm -rf ./client.*
  rm -rf ./ca.*
  rm -rf ./server.*
}

main() {
  genServerCrt
  genClientCert
  creteP12
}
main
