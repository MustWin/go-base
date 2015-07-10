package base

import(
  "net/http"
)

type EndPoint interface {
  GetRoute() string
  GetMethod() string
  GetDescription() string
  GetParameters() interface{}
  GetBody() interface{}
  GetResponses() []Response
  GetHandler() func(http.ResponseWriter, *http.Request)
}

type BaseEndPoint struct {
  Route string
  Method string
  Description string
  Parameters interface{}
  Body interface{}
  Responses []Response
  Handler func(http.ResponseWriter, *http.Request)
}

func (ref *BaseEndPoint) GetRoute() string {
  return ref.Route
}
func (ref *BaseEndPoint) GetMethod() string {
  return ref.Method
}
func (ref *BaseEndPoint) GetDescription() string {
  return ref.Description
}
func (ref *BaseEndPoint) GetParameters() interface{} {
  return ref.GetParameters
}
func (ref *BaseEndPoint) GetBody() interface{} {
  return ref.Body
}
func (ref *BaseEndPoint) GetResponses() []Response{
  return ref.Responses
}
func (ref *BaseEndPoint) GetHandler() func(w http.ResponseWriter, r *http.Request) {
  return ref.Handler
}

var EndPoints []EndPoint
func AddEndPoint(endpoint EndPoint) {
  // TODO: add middlewares

  EndPoints = append(EndPoints, endpoint)
}


