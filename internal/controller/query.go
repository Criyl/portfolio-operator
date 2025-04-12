package controller

import (
	"context"
	"errors"
	"os"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	"carroll.codes/portfolio-operator/internal/config"
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type Querier struct {
	ctx context.Context
	cl  client.Client
}

func (svc *Querier) Init() {
	var (
		restConfig *rest.Config
		err        error
	)

	if _, err := os.Stat(config.Instance.KUBECONFIG); errors.Is(err, os.ErrNotExist) { // if kube config doesn't exist, try incluster config
		restConfig, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
	} else {
		restConfig, err = clientcmd.BuildConfigFromFlags("", config.Instance.KUBECONFIG)
		if err != nil {
			panic(err.Error())
		}
	}

	ctrl.SetLogger(zap.New())

	scheme = runtime.NewScheme()
	utilruntime.Must(opv1.AddToScheme(scheme))
	utilruntime.Must(v1.AddToScheme(scheme))

	svc.ctx = context.Background()
	svc.cl, err = client.New(restConfig, client.Options{Scheme: scheme})
	if err != nil {
		panic(err.Error())
	}

}

func (svc *Querier) ClearPortfolios() {
	portfolios, err := svc.ListPortfolios()
	if err != nil {
		return
	}

	for _, pf := range portfolios.Items {
		svc.cl.Delete(svc.ctx, &pf)
	}
}

func (svc *Querier) ListPortfolios() (*opv1.PortfolioList, error) {
	var portfolios opv1.PortfolioList

	listOpts := []client.ListOption{
		client.InNamespace(
			"",
		),
	}

	err := svc.cl.List(context.Background(), &portfolios, listOpts...)
	if err != nil {
		return nil, err
	}

	return &portfolios, nil
}

func (svc *Querier) ListPortfoliosByTag(tag string) (*opv1.PortfolioList, error) {
	pfList, err := svc.ListPortfolios()
	if err != nil {
		return nil, err
	}

	accepted := []opv1.Portfolio{}

	for _, item := range pfList.Items {
		for _, portfolioTag := range item.Spec.Tags {
			if tag == portfolioTag {
				accepted = append(accepted, item)
			}
		}
	}
	pfList.Items = accepted
	return pfList, nil
}
