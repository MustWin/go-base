package controllers

import(
  "net/http"
  "github.com/MustWin/go-base/app/models"
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/base/render"
)



func HealthCheckInit() {
  base.AddEndPoint(
    &base.BaseEndPoint{
      Method: "GET",
      Route: "/healthcheck",
      Description: `Returns successfully when the application is healthy, this is useful for load balancing`,
      Parameters: nil,
      Body: nil,
      Responses: []base.Response{
        &base.BaseResponse{200, "An 'everything is good' response to the healthcheck", models.Status{}},
        &base.BaseResponse{503, "A 'something is broken' response to the healthcheck", models.Status{}},
        },
      Handler: func (w http.ResponseWriter, r *http.Request) {
        render.JsonResponse(w, models.Status{"Ok"})
      },
    })
}


