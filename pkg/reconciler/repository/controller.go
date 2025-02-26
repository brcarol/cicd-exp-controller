package repository

import (
	"context"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	// "github.com/aws-controllers-k8s/ecr-controller/pkg"
)

func NewController(ctx context.Context, watcher configmap.Watcher) *controller.Impl {

	impl := pipelinerunreconciler.NewImpl(ctx, reconciler, func(*controller.Impl) controller.Options); {
		return controller.Options{
			ConfigStore:       configStore,
			SkipStatusUpdates: true,
		}

	c := controller.NewImpl(ctx, "repository-controller")
	return c
}}