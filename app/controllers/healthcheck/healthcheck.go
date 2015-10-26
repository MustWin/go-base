package healthcheck

import (
	"github.com/MustWin/go-base/app/models"
	"github.com/MustWin/go-base/base"
	"github.com/MustWin/go-base/base/render"
	log "github.com/Sirupsen/logrus"
	httptransport "github.com/go-kit/kit/transport/http"
	"golang.org/x/net/context"
	"net/http"
)

func Endpoints() []base.Endpoint {
	server := httptransport.NewServer(
		context.Background(),
		func(ctx context.Context, req interface{}) (interface{}, error) {
			log.Info("status", "OK")
			return models.Status{"Ok"}, nil
		},
		func(*http.Request) (interface{}, error) {
			return nil, nil
		},
		func(w http.ResponseWriter, status interface{}) error {
			render.JsonResponse(w, status)
			return nil
		},
	)
	return []base.Endpoint{
		&base.BaseEndpoint{
			Method:      "GET",
			Route:       "/healthcheck",
			Description: `Returns successfully when the application is healthy, this is useful for load balancing`,
			Parameters:  nil,
			Body:        nil,
			Responses: []base.Response{
				&base.BaseResponse{200, "An 'everything is good' response", models.Status{}},
				&base.BaseResponse{503, "A 'something is broken' response", models.Status{}},
			},
			Handler: server,
		},
	}
}
