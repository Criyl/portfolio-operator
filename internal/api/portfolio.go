package api

import (
	"net/http"

	opv1 "carroll.codes/portfolio-operator/api/v1"
	_ "carroll.codes/portfolio-operator/docs"
	"carroll.codes/portfolio-operator/internal/controller"
	"github.com/gin-gonic/gin"
)

var querier = controller.Querier{}

// ListEntries return list of all entries
// @Summary return list of all entries
// @Description return list of all entries
// @Tags Portfolio
// @Success 200 {object}  []opv1.PortfolioSpec
// @Router /api/v1/portfolio [get]
func ListPortfolios(c *gin.Context) {
	var (
		list *opv1.PortfolioList
		err  error
	)
	list, err = querier.ListPortfolios()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, portfolioListToListofPortfolioSpec(list))
}

// ListEntriesByTag return list of all entries with a specified tag
// @Summary return list of all entries with a specified tag
// @Description return list of all entries with a specified tag
// @Tags Portfolio
// @Param tag path string true "Tag"
// @Success 200 {object}  []opv1.PortfolioSpec
// @Router /api/v1/portfolio/tag/{tag} [get]
func ListPortfoliosByTag(c *gin.Context) {
	var (
		list *opv1.PortfolioList
		err  error
	)
	list, err = querier.ListPortfoliosByTag(c.Param("tag"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, portfolioListToListofPortfolioSpec(list))
}
