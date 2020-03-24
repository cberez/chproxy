package chproxy

import (
	"fmt"
	"testing"
	"time"

	"net/http"

	"github.com/cberez/chproxy/proxy"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	timeout := 2
	addr := "localhost:8081"
	apiKey := "correct api key"
	// tests expect a chrome headless to be running on port 9222
	addresses := []string{"localhost:9222"}

	go func() {
		p := chproxy.Proxy{ApiKey: apiKey, Timeout: timeout, Addresses: addresses}
		p.ServeAndHandle(addr)
	}()

	t.Run("200 with expected content on correct api key", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", addr), nil)
		req.Header.Set("Api-Key", apiKey)

		client := http.Client{Timeout: time.Duration(time.Duration(timeout) * time.Second)}
		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("error executing request: %v", err)
		}
		defer res.Body.Close()

		sc := res.StatusCode

		assert.Equal(t, http.StatusOK, sc)
	})

	t.Run("401 on incorrect api key", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", addr), nil)
		req.Header.Set("Api-Key", "wrong api key")

		client := http.Client{Timeout: time.Duration(time.Duration(timeout) * time.Second)}
		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("error executing request: %v", err)
		}
		defer res.Body.Close()

		sc := res.StatusCode

		assert.Equal(t, http.StatusUnauthorized, sc)
	})

	t.Run("401 on missing api key", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s", addr), nil)

		client := http.Client{Timeout: time.Duration(time.Duration(timeout) * time.Second)}
		res, err := client.Do(req)
		if err != nil {
			fmt.Printf("error executing request: %v", err)
		}
		defer res.Body.Close()

		sc := res.StatusCode

		assert.Equal(t, http.StatusUnauthorized, sc)
	})
}
