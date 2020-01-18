package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// given
	router := setupRouter()
	w := httptest.NewRecorder()

	// when
	req, _ := http.NewRequest(http.MethodGet, "/health-check", nil)
	router.ServeHTTP(w, req)

	// then
	// assert status code
	assert.Equal(t, 200, w.Code)

	// assert response body
	res := Response{}
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, "unknown", res.Data.Name)
	assert.NotEmpty(t, res.Ts)
}

func TestGetUserName(t *testing.T) {
	// given
	router := setupRouter()
	w := httptest.NewRecorder()

	// when
	name := "vison"
	req, _ := http.NewRequest(http.MethodGet, "/user/"+name, nil)
	router.ServeHTTP(w, req)

	// then
	// assert status code
	assert.Equal(t, 200, w.Code)

	// assert response body
	res := Response{}
	json.Unmarshal([]byte(w.Body.String()), &res)
	assert.Equal(t, name, res.Data.Name)
}
