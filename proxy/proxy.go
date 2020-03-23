package chproxy

import (
	"fmt"
	"log"
	"net/http"
)

type Proxy struct {
	ApiKey string
}

func (p Proxy) Handler(w http.ResponseWriter, r *http.Request) {
	request_key := r.Header.Get("Api-Key")

	if len(request_key) == 0 {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	} else if request_key != p.ApiKey {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	} else {
		log.Printf("api key: '%s'", request_key)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "api key: '%s'", request_key)
	}
}
