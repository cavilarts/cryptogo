package api

import (
	"encoding/json"
	"net/http"

	"cavilarts.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.GetPageInfo())
}