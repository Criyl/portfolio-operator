package controller

import (
	"context"
	"fmt"
	"log"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
)

type ingressReconciler struct {
	client.Client
	scheme     *runtime.Scheme
	kubeClient *kubernetes.Clientset
}

func (r *ingressReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

	var (
		portfolio *opv1.Portfolio
		ingress   v1.Ingress
	)

	// Get the ingress being reconciled
	err := r.Client.Get(ctx, req.NamespacedName, &ingress)

	portfolio = portfolioFromIngress(ingress)
	portfolio.Namespace = req.Namespace

	if !portfolio.IsValid() {
		return ctrl.Result{}, nil
	}
	portfolioName := portfolio.Name

	if err != nil {
		if k8serrors.IsNotFound(err) { // ingress not found, we can clean up resources

			// Find associated Portfolio
			err = r.Client.Get(ctx,
				client.ObjectKey{
					Name:      portfolioName,
					Namespace: req.Namespace,
				},
				portfolio,
			)
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("couldn't find Portfolio %s: %s", portfolioName, err)
			}

			// Delete it
			err = r.Client.Delete(ctx, portfolio, &client.DeleteOptions{})
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("couldn't delete Portfolio %s: %s", portfolioName, err)
			}

			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// See if portfolio already exists
	err = r.Client.Get(ctx, client.ObjectKey{
		Name:      portfolioName,
		Namespace: req.Namespace,
	}, portfolio)
	if err != nil {
		if k8serrors.IsNotFound(err) { // portfolio not found, create one
			err = controllerutil.SetControllerReference(&ingress, portfolio, r.scheme)
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("couldn't set controller reference for Portfolio %s: %s", portfolioName, err)
			}

			err = r.Client.Create(
				ctx,
				portfolio,
			)
			if err != nil {
				return ctrl.Result{}, fmt.Errorf("failed to create Portfolio %s: %s", portfolioName, err)
			}

			log.Println("Created Portfolio ", portfolio.Namespace, "/", portfolio.Name)
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Otherwise update existing Portfolio
	err = r.Client.Update(ctx, portfolio)
	if err != nil {
		return ctrl.Result{}, fmt.Errorf("couldn't update Portfolio %s: %s", portfolioName, err)
	}

	log.Println("Ingress " + portfolioName + " is up-to-date")
	return ctrl.Result{}, nil
}
