package api

import (
	"log"
	"net/http"

	_ "carroll.codes/portfolio-operator/api/v1"
	"carroll.codes/portfolio-operator/internal/config"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Portfolio Operator
// @version 1
// @Description Manage your portfolio natively in your kubernetes cluster.

// @contact.name Christopher Carroll
// @contact.url https://carroll.codes
// @contact.email chris@carroll.codes

// @BasePath /
func MainLoop() {
	querier.Init()
	r := gin.New()

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{
			"/health",
		},
	}))

	r.Use(gin.Recovery())

	r.GET("/health", Health)

	v1 := r.Group("/api/v1")

	entry := v1.Group("/portfolio")
	{
		entry.GET("/", ListPortfolios)
		entry.GET("/tag/:tag", ListPortfoliosByTag)
	}

	if config.Instance.DEBUG == true {
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/swagger/index.html")
		})
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	err := r.Run(config.Instance.HOST + ":" + config.Instance.API_PORT)
	if err != nil {
		log.Fatal(err)
	}
}
