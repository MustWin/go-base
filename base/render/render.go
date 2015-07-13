package render

import (
  "net/http"
  "encoding/json"
)

func JsonResponse(w http.ResponseWriter, responseBodyObj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseBodyObj); err != nil {
	}
}
