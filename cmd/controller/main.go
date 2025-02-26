package main


import (
	"context"
	"knative.dev/pkg/injection"
	"knative.dev/pkg/injection/sharedmain"

	corev1 "k8s.io/api/core/v1"

)

func main() {

	ctx := injection.WithNamespaceScope(context.Background(), corev1.NamespaceAll)

	sharedmain.MainWithContext(
		ctx,
		"controller",
	)
}