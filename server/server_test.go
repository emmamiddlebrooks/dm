package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultFileServer(t *testing.T) {
	t.Run("redirect", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/nonexistent", http.NoBody)
		req.Host = "test"
		rr := httptest.NewRecorder()
		DefaultFileServer(rr, req)
		if rr.Code != http.StatusFound {
			t.Errorf("expected redirect, got status %d", rr.Code)
		}
	})
}
