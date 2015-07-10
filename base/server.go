package base

import (
  "net/http"
  "github.com/gorilla/mux"
  "github.com/spf13/viper"
  "github.com/rs/cors"
  "github.com/codegangsta/negroni"
)


func ServeEndpoints() {
  router := mux.NewRouter()
  for _, endpoint := range EndPoints {
    router.Methods(endpoint.GetMethod()).
           Path(endpoint.GetRoute()).
           HandlerFunc(endpoint.GetHandler())
  }

  cor := cors.New(cors.Options{
    AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
    AllowedMethods:   viper.GetStringSlice("cors.allowed_methods"),
    AllowCredentials: viper.GetBool("cors.allow_credentials"),
  })

  // TODO: make middlewares more pluggable

  server := negroni.New(
    negroni.NewStatic(http.Dir("public")),
    cor,
  )
  server.UseHandler(router)

  // Start the server
  server.Run("0.0.0.0:" + viper.GetString("port"))
}

