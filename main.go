package main

import (
	"fmt"
	"github.com/cberez/chproxy/proxy"
	"strings"
)

func main() {
	addresses := "localhost:9222"
	port := "8080"

	p := chproxy.Proxy{ApiKey: "key", Timeout: 5, Addresses: strings.Split(addresses, ",")}
	p.ServeAndHandle(fmt.Sprintf("localhost:%s", port))
}
