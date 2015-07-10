package app

import(
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/app/controllers/healthcheck"
  log "github.com/Sirupsen/logrus"
)

/*
    This is where you should initialize everything you need for your app to run
*/
func Main() {
  base.Initialize()
  base.RegisterEndpoints(healthcheck.Endpoints()...)

  log.Debug("Registration complete")
  base.ServeEndpoints(base.Endpoints)
}
