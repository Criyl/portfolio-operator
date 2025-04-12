package controller

import (
	"strings"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	v1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func portfolioFromIngress(ingress v1.Ingress) *opv1.Portfolio {
	annotations := ingress.ObjectMeta.Annotations

	tags := strings.Split(annotations["portfolio-operator/tags"], ",")

	return &opv1.Portfolio{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Portfolio",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "portfolio-" + ingress.Namespace + "-" + ingress.Name,
		},
		Spec: opv1.PortfolioSpec{
			Name:        annotations["portfolio-operator/name"],
			Url:         annotations["portfolio-operator/url"],
			Icon:        annotations["portfolio-operator/icon"],
			Blog:        annotations["portfolio-operator/blog"],
			Healthcheck: annotations["portfolio-operator/healthcheck"],
			Tags:        tags,
		},
	}
}
