package main

import (
	"github.com/cberez/chproxy/proxy"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", chproxy.Handler)
	log.Print("listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
