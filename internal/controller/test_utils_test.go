package controller

import (
	"context"
	"testing"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	"github.com/stretchr/testify/assert"
	netv1 "k8s.io/api/networking/v1"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	clientsetscheme "k8s.io/client-go/kubernetes/scheme"

	aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
)

type ReconcilerTester struct {
	t   *testing.T
	ctx *context.Context
	r   *ingressReconciler
}

func (rt *ReconcilerTester) CheckPortfolio(ingress *netv1.Ingress, expectedPortfolio opv1.Portfolio) {
	portfolio := opv1.Portfolio{}

	err := rt.r.Client.Get(*rt.ctx, client.ObjectKey{
		Name:      expectedPortfolio.Name,
		Namespace: expectedPortfolio.Namespace,
	},
		&portfolio,
	)
	assert.NoError(rt.t, err)

	// empty fields for comparison
	portfolio.ResourceVersion = ""
	portfolio.OwnerReferences = nil
	assert.Equal(rt.t, expectedPortfolio, portfolio)
}

func (rt *ReconcilerTester) CheckIngressNotFound(ingress *netv1.Ingress) {
	newPortfolio := opv1.Portfolio{}
	err := rt.r.Client.Get(*rt.ctx, client.ObjectKey{
		Name:      ingress.Name,
		Namespace: ingress.Namespace,
	},
		&newPortfolio,
	)
	if assert.Error(rt.t, err) {
		assert.Equal(rt.t, metav1.StatusReasonNotFound, k8serrors.ReasonForError(err))
	}
}

func (rt *ReconcilerTester) Reconcile(ingress *netv1.Ingress) error {
	recReq := reconcile.Request{
		NamespacedName: types.NamespacedName{
			Name:      ingress.Name,
			Namespace: ingress.Namespace,
		},
	}

	_, err := rt.r.Reconcile(*rt.ctx, recReq)
	return err
}

func (rt *ReconcilerTester) Create(obj client.Object) {
	rt.r.Client.Create(*rt.ctx, obj)
}
func (rt *ReconcilerTester) Update(obj client.Object) {
	rt.r.Client.Update(*rt.ctx, obj)
}
func (rt *ReconcilerTester) Delete(obj client.Object) {
	rt.r.Client.Delete(*rt.ctx, obj)
}

func createFakeReconciler() ingressReconciler {

	restConfig := &rest.Config{}

	kclientset, _ := kubernetes.NewForConfig(restConfig)
	_ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
	_ = opv1.AddToScheme(clientsetscheme.Scheme)
	_ = netv1.AddToScheme(clientsetscheme.Scheme)

	builder := fake.NewClientBuilder().WithScheme(clientsetscheme.Scheme)

	fakeClient := builder.Build()

	return ingressReconciler{
		Client:     fakeClient,
		scheme:     clientsetscheme.Scheme,
		kubeClient: kclientset,
	}
}
