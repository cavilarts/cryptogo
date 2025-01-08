package api

import (
	"encoding/json"
	"net/http"

	"cavilarts.com/go/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var fact []string
		err := json.NewDecoder(r.Body).Decode(&fact)
		println(fact, "HELLO")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data.SetFact(fact)
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}