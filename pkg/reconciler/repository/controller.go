package repository

import (
	"context"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"
	repositoryreconciler "github.com/aws-controllers-k8s/ecr-controller/apis/v1alpha1"

	//  repositoryinformer "github.com/brcarol/cicd-exp-controller/pkg/client/injection/informers/ecr/v1alpha1/repository"
)

func NewController(ctx context.Context, watcher configmap.Watcher) *controller.Impl {
	logger := logging.FromContext(ctx)

	// Add informers

	r := &Reconciler {
		// Add listers and client
	}

	// impl := repositoryreconciler.NewImpl(ctx, r, func(*controller.Impl) controller.Options {
	// 	return controller.Options{
	// 		ConfigStore:       configStore,
	// 		SkipStatusUpdates: true,
	// 	}
	// })

	impl := repositoryreconciler.NewImpl(ctx, r)

	// c := controller.NewImpl(ctx, "repository-controller")
	// return c

	return impl
}