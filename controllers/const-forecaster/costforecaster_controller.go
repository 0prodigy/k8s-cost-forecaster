/*
Copyright 2022 Akash Pathak <pathakvikash9211@gmail.com>.
*/

package constforecaster

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	constforecasterv1alpha1 "github.com/0prodigy/k8s-cost-forecaster/apis/const-forecaster/v1alpha1"
)

// CostForecasterReconciler reconciles a CostForecaster object
type CostForecasterReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=const-forecaster.k8s-cf.io,resources=costforecasters,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=const-forecaster.k8s-cf.io,resources=costforecasters/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=const-forecaster.k8s-cf.io,resources=costforecasters/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the CostForecaster object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *CostForecasterReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *CostForecasterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&constforecasterv1alpha1.CostForecaster{}).
		Complete(r)
}
