package test

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/MustWin/go-base/base"
)

func InitTest() {
	base.Initialize()
}

var testServer http.Handler

func SetServer(endpoints ...base.Endpoint) {
	testServer = base.BuildServer([]http.Handler{}, endpoints)
}

func MakeRequest(method string, url string, stringPayload string) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, url, strings.NewReader(stringPayload))
	if err != nil {
		panic("Failed to create " + url + " REQ")
	}
	req.Header.Set("content-type", "application/json")
	resp := httptest.NewRecorder()
	testServer.ServeHTTP(resp, req)
	return resp
}
