package controller

import (
	"context"
	"testing"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

var (
	gatewayNamespaceFixture v1.Namespace = v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "portfolio",
		},
	}
	gatewayNamespacedFixture types.NamespacedName = types.NamespacedName{
		Namespace: gatewayNamespaceFixture.Name,
		Name:      "gateway-httproute",
	}

	gatewayFixture gatewayv1.HTTPRoute = gatewayv1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gatewayNamespacedFixture.Name,
			Namespace: gatewayNamespacedFixture.Namespace,
			Annotations: map[string]string{
				"portfolio-operator/name": "portfolio",
				"portfolio-operator/url":  "portfolio.com",
			},
		},
	}
)

func TestReconciler_DeletedGatewayNoPortfolio(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler[gatewayv1.HTTPRoute]()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	gateway := gatewayFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromObject(gateway)

	rt.Create(&namespaceFixture)
	rt.Create(gateway)

	err := rt.Reconcile(gateway)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(gateway, *expectedPortfolio)

	// Test Reconcile Deleted owner gateway and portfolio
	rt.Delete(gateway)
	rt.Delete(expectedPortfolio)

	err = rt.Reconcile(gateway)
	assert.Error(rt.t, err)

	rt.CheckObjectNotFound(gateway)
}

func TestReconciler_ReconcileCreateMutateGateway(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler[gatewayv1.HTTPRoute]()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	gateway := gatewayFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromObject(gateway)

	rt.Create(&namespaceFixture)
	rt.Create(gateway)
	err := rt.Reconcile(gateway)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(gateway, *expectedPortfolio)

	// Test Reconcile Portfolio with mutated gateway
	newPortfolioName := "Portfolio"
	gateway.Annotations["portfolio-operator/name"] = newPortfolioName
	expectedPortfolio.Spec.Name = newPortfolioName

	rt.Update(gateway)
	err = rt.Reconcile(gateway)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(gateway, *expectedPortfolio)
}

func TestReconciler_ReconcileCreateDeleteGateway(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler[gatewayv1.HTTPRoute]()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}

	gateway := gatewayFixture.DeepCopy()
	expectedPortfolio := portfolioCreateFromObject(gateway)

	rt.Create(&namespaceFixture)
	rt.Create(gateway)
	err := rt.Reconcile(gateway)
	assert.NoError(rt.t, err)

	rt.CheckPortfolio(gateway, *expectedPortfolio)

	// Test Reconcile Deleted owner gateway
	rt.Delete(gateway)
	err = rt.Reconcile(gateway)
	assert.NoError(rt.t, err)

	rt.CheckObjectNotFound(gateway)
}

func TestReconciler_ReconcileInvalidGateway(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reconciler := createFakeReconciler[gatewayv1.HTTPRoute]()

	rt := ReconcilerTester{
		t:   t,
		ctx: &ctx,
		r:   &reconciler,
	}
	testgateway := gatewayv1.HTTPRoute{
		ObjectMeta: metav1.ObjectMeta{
			Name:      gatewayNamespacedFixture.Name,
			Namespace: gatewayNamespacedFixture.Namespace,
		},
	}
	rt.Create(&testgateway)

	err := rt.Reconcile(&testgateway)
	assert.NoError(rt.t, err)

	portfolioList := opv1.PortfolioList{}
	rt.r.Client.List(*rt.ctx, &portfolioList)
	assert.Equal(t, []opv1.Portfolio{}, portfolioList.Items)
}
