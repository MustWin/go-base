package healthcheck

import(
  "net/http"
  "golang.org/x/net/context"
  "github.com/MustWin/go-base/app/models"
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/base/render"
  httptransport "github.com/go-kit/kit/transport/http"
)



func Endpoints() []base.Endpoint {
  return []base.Endpoint{
    &base.BaseEndpoint{
      Method: "GET",
      Route: "/healthcheck",
      Description: `Returns successfully when the application is healthy, this is useful for load balancing`,
      Parameters: nil,
      Body: nil,
      Responses: []base.Response{
        &base.BaseResponse{200, "An 'everything is good' response", models.Status{}},
        &base.BaseResponse{503, "A 'something is broken' response", models.Status{}},
        },
      Handler: httptransport.Server{
		  Context:    context.Background(),
		  DecodeFunc: func(*http.Request) (interface{}, error) {
            return nil, nil
          },
		  Endpoint: func(ctx context.Context, req interface{}) (interface{}, error) {
            return models.Status{"Ok"}, nil
          },
		  EncodeFunc: func(w http.ResponseWriter, status interface{}) error {
            render.JsonResponse(w, status)
            return nil
          },
	    },
    },
  }
}

