package repository

import (
	"context"

	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	repositoryreconciler "cicd-exp-controller/pkg/client/injection/reconciler/apis/v1alpha1/repository"
)

func NewController(ctx context.Context, watcher configmap.Watcher) *controller.Impl {
	//logger := logging.FromContext(ctx)

	// Add informers

	reconciler := &RepositoryReconciler{
		//Add listers and client
	}

	impl := repositoryreconciler.NewImpl(ctx, reconciler, func(*controller.Impl) controller.Options {
		return controller.Options{
			//ConfigStore:       configStore,
			SkipStatusUpdates: true,
		}
	})

	return impl
}
