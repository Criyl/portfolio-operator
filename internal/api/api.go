package api

import (
	"log"
	"net/http"
	"os"

	_ "carroll.codes/portfolio-operator/api/v1"
	"carroll.codes/portfolio-operator/internal/config"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	clientID     = os.Getenv("OIDC_CLIENT_ID")
	clientSecret = os.Getenv("OIDC_CLIENT_SECRET")
	redirectURL  = os.Getenv("OIDC_REDIRECT_URL")
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
	r := gin.Default()

	if config.Instance.DEBUG == true {
		r.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusFound, "/swagger/index.html")
		})
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	r.GET("/health", Health)

	v1 := r.Group("/api/v1")

	entry := v1.Group("/portfolio")
	{
		entry.GET("/", ListPortfolios)
		entry.GET("/tag/:tag", ListPortfoliosByTag)
	}

	err := r.Run(config.Instance.HOST + ":" + config.Instance.PORT)
	if err != nil {
		log.Fatal(err)
	}
}
