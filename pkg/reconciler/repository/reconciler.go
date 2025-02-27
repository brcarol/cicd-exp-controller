package repository

import (
	"context"

	"knative.dev/pkg/reconciler"

	v1alpha1 "github.com/aws-controllers-k8s/ecr-controller/apis/v1alpha1"
)

type RepositoryReconciler struct {
}

// ReconcileKind implements Interface.ReconcileKind.
func (r *RepositoryReconciler) ReconcileKind(ctx context.Context, repository *v1alpha1.Repository) reconciler.Event {
	return nil
}
