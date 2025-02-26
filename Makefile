lint:
	@golangci-lint run

verify-codegen:
	@./hack/verify-codegen.sh
