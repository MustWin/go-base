package app

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/carbocation/interpose/adaptors"
	"github.com/codegangsta/negroni"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"github.com/unrolled/secure"

	"github.com/MustWin/go-base/base"
)

// Main is where you should initialize everything you need for your app to run
func Main() {
	base.Initialize()

	log.Debug("Registration complete")
	base.ServeEndpoints(setupMiddleware(), base.Endpoints())
}

func setupMiddleware() []http.Handler {
	return []http.Handler{
		adaptors.HandlerFromNegroni(cors.New(cors.Options{
			AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
			AllowedMethods:   viper.GetStringSlice("cors.allowed_methods"),
			AllowCredentials: viper.GetBool("cors.allow_credentials"),
		})),
		adaptors.HandlerFromNegroni(negroni.HandlerFunc(secure.New(secure.Options{
			AllowedHosts:            []string{},
			SSLRedirect:             false,
			SSLTemporaryRedirect:    false,
			SSLHost:                 "",
			SSLProxyHeaders:         map[string]string{},
			STSSeconds:              0,
			STSIncludeSubdomains:    false,
			STSPreload:              false,
			ForceSTSHeader:          false,
			FrameDeny:               false,
			CustomFrameOptionsValue: "",
			ContentTypeNosniff:      false,
			BrowserXssFilter:        false,
			ContentSecurityPolicy:   "",
			IsDevelopment:           viper.GetString("environment") == "development",
		}).HandlerFuncWithNext)),
	}
}
