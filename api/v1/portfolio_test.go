package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortfolioValidation_Valid(t *testing.T) {
	validPortfolio := Portfolio{
		Spec: PortfolioSpec{
			Name: "portfolio",
			Url:  "portfolio.com",
		},
	}

	assert.True(t, validPortfolio.IsValid())
}
func TestPortfolioValidation_Empty(t *testing.T) {
	invalidPortfolio := Portfolio{}

	assert.False(t, invalidPortfolio.IsValid())
}
