package repository

import (
	"context"

	"knative.dev/pkg/logging"
	"knative.dev/pkg/reconciler"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	repositoryclientset "cicd-exp-controller/pkg/client/clientset/versioned"

	v1alpha1 "github.com/aws-controllers-k8s/ecr-controller/apis/v1alpha1"
)

type RepositoryReconciler struct {
	client repositoryclientset.Interface
}

// ReconcileKind implements Interface.ReconcileKind.
func (r *RepositoryReconciler) ReconcileKind(ctx context.Context, repository *v1alpha1.Repository) reconciler.Event {
	logger := logging.FromContext(ctx)

	logger.Infof("Reconciling Repository %s", repository.Name)

	if value, ok := repository.GetAnnotations()["reconciled"]; ok {
		logger.Infof("Repository %s already have the reconciled annotation with value %s", repository.Name, value)
		return nil
	}

	annotations := repository.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}

	annotations["reconciled"] = "true"
	repository.SetAnnotations(annotations)

	r.client.EcrV1alpha1().Repositories(repository.Namespace).Update(ctx, repository, metav1.UpdateOptions{})

	return nil
}
