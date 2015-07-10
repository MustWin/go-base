package healthcheck

import (
  "net/http"
  "os"
  "testing"
  "github.com/MustWin/go-base/test"
)

func TestMain(m *testing.M) {
	test.InitTest()
	test.SetServer(Endpoints()...)
	os.Exit(m.Run())
}

func TestGetHealthCheck(t *testing.T) {
  resp := test.MakeRequest("GET", "/healthcheck", "")
  if resp.Code != http.StatusOK{
  	t.Fatalf("Invalid response code (%d) for GET to /healthcheck. Expected %d", resp.Code, http.StatusOK)
  }
}
