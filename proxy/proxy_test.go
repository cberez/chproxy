package chproxy

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/cberez/chproxy/proxy"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	t.Run("200 on existing api key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		request.Header.Set("Api-Key", "some api key")
		response := httptest.NewRecorder()

		chproxy.Handler(response, request)
		sc := response.Result().StatusCode

		assert.Equal(t, http.StatusOK, sc)
	})

	t.Run("401 on missing api key", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		chproxy.Handler(response, request)
		sc := response.Result().StatusCode

		assert.Equal(t, http.StatusUnauthorized, sc)
	})
}
