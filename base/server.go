package base

import (
	"log"
	"net/http"
	"time"

	"github.com/carbocation/interpose"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// BuildServer applies context-less middlewares. Ones with context will require some more work
func BuildServer(middlewares []http.Handler, endpoints []Endpoint) http.Handler {
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

// ServeEndpoints starts the server with the configured middleware and
// endpoints.
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
