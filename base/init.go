package base

import (
  "math/rand"
  "os"
  "time"
  "github.com/spf13/viper"
  log "github.com/Sirupsen/logrus"
)



func Initialize() {
  rand.Seed(time.Now().UTC().UnixNano())
  LoadConfig()

  log.SetFormatter(&log.TextFormatter{})
  log.SetOutput(os.Stderr)

  level, err := log.ParseLevel(viper.GetString("log_level"))
  if err != nil {
    panic("Invalid log level setting")
  }
  log.SetLevel(level)
}
