package chproxy

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	apiKey := r.Header.Get("Api-Key")

	if len(apiKey) > 0 {
		log.Printf("api key: '%s'", apiKey)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "api key: '%s'", apiKey)
	} else {
		log.Print("missing api key")
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
