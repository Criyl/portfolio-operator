package api

// import (
// 	"context"
// 	"log"
// 	"net/http"

// 	_ "carroll.codes/portfolio-operator/docs"
// 	"carroll.codes/portfolio-operator/internal/store"
// 	"github.com/coreos/go-oidc/v3/oidc"
// 	"github.com/gin-gonic/gin"
// 	"golang.org/x/oauth2"
// )

// type OIDCContext struct {
// 	Config   *oidc.Config
// 	Provider *oidc.Provider
// 	Verifier *oidc.IDTokenVerifier
// }

// func (ctx *OIDCContext) init() {

// 	provider, err := oidc.NewProvider(context.Background(), "https://accounts.google.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	oidcConfig := oauth2.Config{
// 		ClientID:     clientID,
// 		ClientSecret: clientSecret,
// 		RedirectURL:  redirectURL,
// 		Endpoint:     provider.Endpoint(),
// 		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
// 	}

// 	verifier := provider.Verifier(*oidcConfig)

// 	ctx.Config = *oidcConfig
// 	ctx.Provider = provider
// 	ctx.Verifier = verifier
// }

// type callbackResponse struct {
// 	Data []store.Entry `json:"data"`
// }

// // OIDCLogin return redirect user to authentication provider
// // @Summary return list of all
// // @Description return list of all entries from the database
// // @Tags Auth
// // @Success 200 {object} entryResponse
// // @Router /login [get]
// func OIDCLogin(c *gin.Context) {
// 	entries := store.EntryStoreInstance.GetEntries()

// 	c.JSON(http.StatusOK, entryResponse{Data: entries})
// }

// // OIDCCallback return Handle authentication
// // @Summary return list of all
// // @Description return list of all entries from the database
// // @Tags Auth
// // @Success 200 {object} entryResponse
// // @Router /callback [get]
// func OIDCCallback(c *gin.Context) {
// 	entries := store.EntryStoreInstance.GetEntries()

// 	c.JSON(http.StatusOK, entryResponse{Data: entries})
// }
