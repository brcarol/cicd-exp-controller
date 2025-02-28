package repository

import (
	"context"
	"time"

	"k8s.io/client-go/tools/cache"
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	repositoryclient "cicd-exp-controller/pkg/client/injection/client"
	repositoryinformer "cicd-exp-controller/pkg/client/injection/informers/apis/v1alpha1/repository"
	repositoryreconciler "cicd-exp-controller/pkg/client/injection/reconciler/apis/v1alpha1/repository"
)

func NewController(ctx context.Context, watcher configmap.Watcher) *controller.Impl {
	logger := logging.FromContext(ctx)

	// configStore := config.NewStore(logger.Named("configs"))
	// configStore.WatchConfigs(watcher)

	// Add informers
	informer := repositoryinformer.Get(ctx)

	reconciler := &RepositoryReconciler{
		client: repositoryclient.Get(ctx),
	}

	logger = logger.Named("cicd-exp-controller")
	impl := repositoryreconciler.NewImpl(ctx, reconciler, func(*controller.Impl) controller.Options {
		return controller.Options{
			SkipStatusUpdates: true,
		}
	})

	logger.Info("Setting up informer")
	informer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: impl.Enqueue,
		UpdateFunc: func(_, newObj interface{}) {
			impl.EnqueueAfter(newObj, 10*time.Second)
		},
	})

	return impl
}
