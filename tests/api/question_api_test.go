package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gobot/api/handlers"
)

func TestLogin(t *testing.T) {
	payload := []byte(`{"username": "testuser", "password": "password"}`)
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.Login)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, http.StatusOK)
	}

	var token string
	err = json.Unmarshal(rr.Body.Bytes(), &token)
	if err != nil {
		t.Fatal(err)
	}

	if token == "" {
		t.Errorf("handler returned empty token")
	}
}

