package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/cberez/chproxy/proxy"
)

func main() {

	var addresses string
	flag.StringVar(&addresses, "addresses", "localhost:9222", "comma separated list of chrome headless addresses")

	var port string
	flag.StringVar(&port, "port", "8080", "port to listen to for incoming connections")

	var timeout int
	flag.IntVar(&timeout, "timeouyt", 5, "connection timeout in seconds")

	var apiKey string
	flag.StringVar(&apiKey, "apikey", "key", "api key to use")

	flag.Parse()

	p := chproxy.Proxy{ApiKey: apiKey, Timeout: timeout, Addresses: strings.Split(addresses, ",")}

	log.Printf("listening on port %s", port)

	p.ServeAndHandle(fmt.Sprintf("localhost:%s", port))
}
