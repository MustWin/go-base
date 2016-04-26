package base

import (
	"fmt"
	"math/rand"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Initialize the application, including logging and config
func Initialize() {
	rand.Seed(time.Now().UTC().UnixNano())
	loadConfig()

	log.SetFormatter(&log.TextFormatter{})

	ll := viper.GetString("log_level")
	if ll == "" {
		ll = "debug"
	}
	level, err := log.ParseLevel(ll)
	if err != nil {
		panic(fmt.Sprintf("Invalid log level setting: %v", err))
	}
	log.SetLevel(level)
}
