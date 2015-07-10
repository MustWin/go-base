package app

import(
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/app/controllers/healthcheck"
)

/*
    This is where you should initialize everything you need for your app to run
*/
func Main() {
  base.RegisterEndpoints(healthcheck.Endpoints()...)

  base.ServeEndpoints()
}
