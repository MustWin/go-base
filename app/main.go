package app

import(
  "fmt"
  "github.com/MustWin/go-base/base"
  "github.com/MustWin/go-base/app/controllers"
)

/*
    This is where you should initialize everything you need for your app to run
*/
func Main() {
  fmt.Println("HERE")
  controllers.HealthCheckInit()

  base.ServeEndpoints()
}
