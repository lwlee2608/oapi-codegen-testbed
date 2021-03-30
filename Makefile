.PHONY: generate
generate:
	oapi-codegen spec.yaml > ./generated.go

