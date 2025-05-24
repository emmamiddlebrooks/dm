package guest

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGuestHandler(t *testing.T) {
	t.Run("POST", func(t *testing.T) {
		guestString := `{"first_name":"John","last_name":"Doe","email":"johndoe@test.com"}`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		handler := NewGuestHandler(context.Background(),
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	t.Run("DELETE", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/submit", http.NoBody)
		rr := httptest.NewRecorder()
		handler := NewGuestHandler(context.Background(),
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusMethodNotAllowed)
		assert.Contains(t, rr.Body.String(), "invalid method")
	})
}
