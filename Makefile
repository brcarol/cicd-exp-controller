lint:
	@golangci-lint run

verify-codegen:
	@./hack/verify-codegen.sh

install-dev:
	@export KIND_CLUSTER_NAME=cicd-exp && export KO_DOCKER_REPO=kind.local && kustomize build config/base | ko apply --platform linux/arm64 -f -

