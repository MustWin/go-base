package base

import (
  "math/rand"
  "time"
)

func Initialize() {
  rand.Seed(time.Now().UTC().UnixNano())
  LoadConfig()
}
