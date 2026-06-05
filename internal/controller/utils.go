package controller

import (
	"strings"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type GenericObject struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

func generatePortfolioName(ingressNamespacedName types.NamespacedName) string {
	return "portfolio-" + ingressNamespacedName.Namespace + "-" + ingressNamespacedName.Name
}

func portfolioCreateFromObject(object client.Object) *opv1.Portfolio {
	portfolio := opv1.Portfolio{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Portfolio",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: generatePortfolioName(types.NamespacedName{
				Name:      object.GetName(),
				Namespace: object.GetNamespace(),
			}),
			Namespace: object.GetNamespace(),
		},
	}
	return portfolioUpdateFromObject(object, &portfolio)
}

func portfolioUpdateFromObject(object client.Object, portfolio *opv1.Portfolio) *opv1.Portfolio {
	annotations := object.GetAnnotations()

	tags := strings.Split(annotations["portfolio-operator/tags"], ",")

	portfolio.Spec = opv1.PortfolioSpec{
		Name:        annotations["portfolio-operator/name"],
		Url:         annotations["portfolio-operator/url"],
		Icon:        annotations["portfolio-operator/icon"],
		Blog:        annotations["portfolio-operator/blog"],
		Healthcheck: annotations["portfolio-operator/healthcheck"],
		Tags:        tags,
	}

	return portfolio
}
