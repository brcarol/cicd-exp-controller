lint:
	@golangci-lint run

verify-codegen:
	@./hack/verify-codegen.sh

install-dev:
	@export KO_DOCKER_REPO=localhost:5001 && kustomize build config/base | ko resolve --platform linux/arm64 -f - > ./setup/install-manifests.yaml
