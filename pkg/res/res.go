package res

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, ststusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ststusCode)
	json.NewEncoder(w).Encode(data)
}
