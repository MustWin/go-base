package base

import(
  httptransport "github.com/go-kit/kit/transport/http"
)

type Endpoint interface {
  GetRoute() string
  GetMethod() string
  GetDescription() string
  GetParameters() interface{}
  GetBody() interface{}
  GetResponses() []Response
  GetHandler() httptransport.Server
}

type BaseEndpoint struct {
  Route string
  Method string
  Description string
  Parameters interface{}
  Body interface{}
  Responses []Response
  Handler httptransport.Server
}

func (ref *BaseEndpoint) GetRoute() string {
  return ref.Route
}
func (ref *BaseEndpoint) GetMethod() string {
  return ref.Method
}
func (ref *BaseEndpoint) GetDescription() string {
  return ref.Description
}
func (ref *BaseEndpoint) GetParameters() interface{} {
  return ref.GetParameters
}
func (ref *BaseEndpoint) GetBody() interface{} {
  return ref.Body
}
func (ref *BaseEndpoint) GetResponses() []Response{
  return ref.Responses
}
func (ref *BaseEndpoint) GetHandler() httptransport.Server {
  return ref.Handler
}

var Endpoints []Endpoint
func RegisterEndpoints(endpoints ...Endpoint) {
  // TODO: add middlewares
  // TODO: re-add middlewares
  /*cor := cors.New(cors.Options{
    AllowedOrigins:   viper.GetStringSlice("cors.allowed_origins"),
    AllowedMethods:   viper.GetStringSlice("cors.allowed_methods"),
    AllowCredentials: viper.GetBool("cors.allow_credentials"),
  })*/

  // TODO: Prevent duplicate additions

  Endpoints = append(Endpoints, endpoints...)
}


