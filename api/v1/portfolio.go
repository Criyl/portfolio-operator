package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type PortfolioList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Portfolio `json:"items"`
}

type Portfolio struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PortfolioSpec `json:"spec"`
}

type PortfolioSpec struct {
	Name        string   `json:"name"`
	Url         string   `json:"url"`
	Blog        string   `json:"blog"`
	Icon        string   `json:"icon"`
	Healthcheck string   `json:"healthcheck"`
	Tags        []string `json:"tags"`
}

func (Portfolio *Portfolio) IsValid() bool {
	return Portfolio.Spec.Name != "" && Portfolio.Spec.Url != ""
}
