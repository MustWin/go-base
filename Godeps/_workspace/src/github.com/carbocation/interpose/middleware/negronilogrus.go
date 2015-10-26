package middleware

import (
	"net/http"

	"github.com/MustWin/go-base/Godeps/_workspace/src/github.com/carbocation/interpose/adaptors"
	"github.com/meatballhat/negroni-logrus"
)

func NegroniLogrus() func(http.Handler) http.Handler {
	return adaptors.FromNegroni(negronilogrus.NewMiddleware())
}
