package app

import(
  "net/http"
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/app/controllers/healthcheck"
  log "github.com/Sirupsen/logrus"
  "github.com/spf13/viper"

  "github.com/codegangsta/negroni"
  "github.com/carbocation/interpose/adaptors"
  "github.com/rs/cors"
  "github.com/unrolled/secure"
)

/*
    This is where you should initialize everything you need for your app to run
*/
func Main() {
  base.Initialize()

  // Create Endpoints
  base.RegisterEndpoints(healthcheck.Endpoints()...)

  log.Debug("Registration complete")
  base.ServeEndpoints(setupMiddleware(), base.Endpoints)
}

func setupMiddleware() []http.Handler {
  return []http.Handler {
    adaptors.HandlerFromNegroni(cors.New(cors.Options{
      AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
      AllowedMethods:   viper.GetStringSlice("cors.allowed_methods"),
      AllowCredentials: viper.GetBool("cors.allow_credentials"),
    })),
    adaptors.HandlerFromNegroni(negroni.HandlerFunc(secure.New(secure.Options{
      AllowedHosts: []string{},
      SSLRedirect: false,
      SSLTemporaryRedirect: false,
      SSLHost: "",
      SSLProxyHeaders: map[string]string{},
      STSSeconds: 0,
      STSIncludeSubdomains: false,
      STSPreload: false,
      ForceSTSHeader: false,
      FrameDeny: false,
      CustomFrameOptionsValue: "",
      ContentTypeNosniff: false,
      BrowserXssFilter: false,
      ContentSecurityPolicy: "",
      IsDevelopment: viper.GetString("environment") == "development",
    }).HandlerFuncWithNext)),
  }
}
