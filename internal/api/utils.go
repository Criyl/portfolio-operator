package api

import (
	opv1 "carroll.codes/portfolio-operator/api/v1"
)

func portfolioToList(pfList *opv1.PortfolioList) []opv1.PortfolioSpec {
	list := []opv1.PortfolioSpec{}
	for _, item := range pfList.Items {
		list = append(list, item.Spec)
	}
	return list
}