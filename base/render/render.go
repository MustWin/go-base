package render

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse writes a response in JSON format.
func JSONResponse(w http.ResponseWriter, responseBodyObj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseBodyObj); err != nil {
		w.Write([]byte(fmt.Sprintf("{error: %q}", err.Error())))
	}
}
