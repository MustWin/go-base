package base

import (
  "log"
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/spf13/viper"
  //"github.com/rs/cors"
)

func BuildServer(endpoints ...Endpoint) http.Handler {
  router := mux.NewRouter()
  for _, endpoint := range endpoints {
    router.Methods(endpoint.GetMethod()).
           Path(endpoint.GetRoute()).
           HandlerFunc(endpoint.GetHandler().ServeHTTP)
  }
  return router
}

func ServeEndpoints(endpoints []Endpoint) {

  server := &http.Server{
    Addr:           "0.0.0.0:" + viper.GetString("port"),
    Handler:        BuildServer(endpoints...),
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  // Start the server
  log.Fatal(server.ListenAndServe())
}

