# einstein-golang-sdk

Generated with [go-swagger](https://github.com/go-swagger/go-swagger)

How to generate:
```
swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger'
swagger generate client -f ./openapi-2.0.0.spec.yaml -A einstein-golang-sdk  ./client.go
```

Install deps:
```
go get -u -f ./...
```
