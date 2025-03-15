package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloAPI(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/api/hello", nil)
	w := httptest.NewRecorder()

	handler(w, r)
	resp := w.Result()

	_, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("cannot read test response: %v", err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("got = %d, want = 200", resp.StatusCode)
	}
}
