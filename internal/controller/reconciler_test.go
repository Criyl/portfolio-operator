package controller

import (
	"context"
	"testing"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var (
	namespaceFixture v1.Namespace = v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "portfolio",
		},
	}
	ingressNamespacedFixture types.NamespacedName = types.NamespacedName{
		Namespace: namespaceFixture.Name,
		Name:      "ingress",
	}

	ingressFixture netv1.Ingress = netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressNamespacedFixture.Name,
			Namespace: ingressNamespacedFixture.Namespace,
			Annotations: map[string]string{
				"portfolio-operator/name": "portfolio",
				"portfolio-operator/url":  "portfolio.com",
			},
		},
	}
)

func TestReconciler_DeletedIngressNoPortfolio(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	ingress := ingressFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromIngress(*ingress)

	rt.Create(&namespaceFixture)
	rt.Create(ingress)

	err := rt.Reconcile(ingress)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(ingress, *expectedPortfolio)

	// Test Reconcile Deleted owner ingress and portfolio
	rt.Delete(ingress)
	rt.Delete(expectedPortfolio)

	err = rt.Reconcile(ingress)
	assert.Error(rt.t, err)

	rt.CheckIngressNotFound(ingress)
}

func TestReconciler_ReconcileCreateMutate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	ingress := ingressFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromIngress(*ingress)

	rt.Create(&namespaceFixture)
	rt.Create(ingress)
	err := rt.Reconcile(ingress)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(ingress, *expectedPortfolio)

	// Test Reconcile Portfolio with mutated Ingress
	newPortfolioName := "Portfolio"
	ingress.Annotations["portfolio-operator/name"] = newPortfolioName
	expectedPortfolio.Spec.Name = newPortfolioName

	rt.Update(ingress)
	err = rt.Reconcile(ingress)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(ingress, *expectedPortfolio)
}

func TestReconciler_ReconcileCreateDelete(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	ingress := ingressFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromIngress(*ingress)

	rt.Create(&namespaceFixture)
	rt.Create(ingress)
	err := rt.Reconcile(ingress)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(ingress, *expectedPortfolio)

	// Test Reconcile Deleted owner ingress
	rt.Delete(ingress)
	err = rt.Reconcile(ingress)
	assert.NoError(rt.t, err)

	rt.CheckIngressNotFound(ingress)
}

func TestReconciler_ReconcileInvalidIngress(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}
	testIngress := netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressNamespacedFixture.Name,
			Namespace: ingressNamespacedFixture.Namespace,
		},
	}
	rt.Create(&testIngress)

	err := rt.Reconcile(&testIngress)
	assert.NoError(rt.t, err)

	portfolioList := opv1.PortfolioList{}
	rt.r.Client.List(*rt.ctx, &portfolioList)
	assert.Equal(t, []opv1.Portfolio{}, portfolioList.Items)
}
