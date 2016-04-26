package base

import (
	"fmt"

	httptransport "github.com/go-kit/kit/transport/http"
)

// Endpoint describes a generic endpoint defining a service.
type Endpoint interface {
	GetRoute() string
	GetMethod() string
	GetDescription() string
	GetParameters() interface{}
	GetBody() interface{}
	GetResponses() []Response
	GetHandler() *httptransport.Server
}

// DefaultEndpoint is the standard base Endpoint implementation.
type DefaultEndpoint struct {
	Route       string
	Method      string
	Description string
	Parameters  interface{}
	Body        interface{}
	Responses   []Response
	Handler     *httptransport.Server
}

// GetRoute returns the Route
func (ref *DefaultEndpoint) GetRoute() string {
	return ref.Route
}

// GetMethod returns the Method
func (ref *DefaultEndpoint) GetMethod() string {
	return ref.Method
}

// GetDescription returns the Description
func (ref *DefaultEndpoint) GetDescription() string {
	return ref.Description
}

// GetParameters returns the Parameters
func (ref *DefaultEndpoint) GetParameters() interface{} {
	return ref.GetParameters
}

// GetBody returns the Body
func (ref *DefaultEndpoint) GetBody() interface{} {
	return ref.Body
}

// GetResponses returns the slice of Response
func (ref *DefaultEndpoint) GetResponses() []Response {
	return ref.Responses
}

// GetHandler returns the Handler
func (ref *DefaultEndpoint) GetHandler() *httptransport.Server {
	return ref.Handler
}

var registeredEndpoints map[string]Endpoint

// RegisterEndpoints adds the given endpoints to the registry of
// endpoints.
func RegisterEndpoints(endpoints ...Endpoint) {
	for _, endpoint := range endpoints {
		if endpoint == nil {
			continue
		}
		key := fmt.Sprintf("%s-%s", endpoint.GetRoute(), endpoint.GetMethod())
		registeredEndpoints[key] = endpoint
	}
}

// Endpoints returns the list of registered Endpoints
func Endpoints() []Endpoint {
	endpoints := []Endpoint{}
	for _, endpoint := range registeredEndpoints {
		endpoints = append(endpoints, endpoint)
	}
	return endpoints
}
