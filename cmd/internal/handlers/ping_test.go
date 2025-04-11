package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhruvsaxena1998/aio/cmd/internal/helpers"
	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	// Create a request to the ping endpoint
	req, err := http.NewRequest("GET", "/ping", nil)
	assert.NoError(t, err)

	// Create a recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler
	handler := http.HandlerFunc(PingHandler)
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Parse the response
	var response helpers.Response
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Verify the response
	assert.True(t, response.Success)

	// Check that message is "pong"
	data, ok := response.Data.(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, "pong", data["message"])
}
