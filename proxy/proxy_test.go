package chproxy

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/cberez/chproxy/proxy"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("200 on correct api key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Api-Key", "correct api key")
		response := httptest.NewRecorder()

		p := chproxy.Proxy{ApiKey: "correct api key"}
		p.Handler(response, request)

		sc := response.Result().StatusCode

		assert.Equal(t, http.StatusOK, sc)
	})

	t.Run("401 on incorrect api key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Api-Key", "wrong api key")
		response := httptest.NewRecorder()

		p := chproxy.Proxy{ApiKey: "correct api key"}
		p.Handler(response, request)

		sc := response.Result().StatusCode

		assert.Equal(t, http.StatusUnauthorized, sc)
	})

	t.Run("401 on missing api key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		p := chproxy.Proxy{ApiKey: "correct api key"}
		p.Handler(response, request)

		sc := response.Result().StatusCode

		assert.Equal(t, http.StatusUnauthorized, sc)
	})
}
