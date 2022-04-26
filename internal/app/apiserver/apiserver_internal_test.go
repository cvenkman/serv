package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIserver_HandleHello(t *testing.T) {
	server := New(NewConfig())
	recorder := httptest.NewRecorder()
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)

	server.handleHello().ServeHTTP(recorder, request)
	assert.Equal(t, recorder.Body.String(), "Hi")
}