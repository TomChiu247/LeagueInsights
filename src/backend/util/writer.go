package util

import (
	"encoding/json"
	"net/http"
)

// JSON converts maps in JSON
func JSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	bytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json;")
	w.WriteHeader(status)
	w.Write(bytes)
}
