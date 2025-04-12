package api

import (
	"net/http"

	_ "carroll.codes/portfolio-operator/docs"
	"github.com/gin-gonic/gin"
)

// Health return return ok status if service is healthy
// @Summary return return ok status if service is healthy
// @Description return return ok status if service is healthy
// @Tags Health
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
