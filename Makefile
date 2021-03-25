.PHONY: client server generate
client:
	go build -o client cmd/client/*.go
server:
	go build -o server cmd/server/*.go
generate:
	oapi-codegen spec.yaml  > pkg/generated/generated.go

