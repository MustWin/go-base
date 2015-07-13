package base

import (
  "log"
  "time"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/spf13/viper"
  "github.com/carbocation/interpose"
)

type Context struct {
  Test string
}


func BuildServer(middlewares []http.Handler, endpoints []Endpoint) http.Handler { // This function applies context-less middlewares. Ones with context will require some more work
  handlerChain := interpose.New()
  for _, middle := range middlewares {
    handlerChain.UseHandler(middle)
  }

  router := mux.NewRouter()

  for _, endpoint := range endpoints {
    router.Methods(endpoint.GetMethod()).
           Path(endpoint.GetRoute()).
           HandlerFunc(endpoint.GetHandler().ServeHTTP)
  }
  handlerChain.UseHandler(router)
  return handlerChain
}

func ServeEndpoints(middlewares []http.Handler, endpoints []Endpoint) {

  server := &http.Server{
    Addr:           "0.0.0.0:" + viper.GetString("port"),
    Handler:        BuildServer(middlewares, endpoints),
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  // Start the server
  log.Fatal(server.ListenAndServe())
}

