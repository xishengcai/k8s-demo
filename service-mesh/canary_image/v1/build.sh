
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/canary main.go
docker build -t xishengcai/canary ./
docker push xishengcai/canary
