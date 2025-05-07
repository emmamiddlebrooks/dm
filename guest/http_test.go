package guest

import (
	"bytes"
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGuestHandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("POST", func(mt *mtest.T) {
		guestString := `{"first_name":"John","last_name":"Doe","email":"johndoe@test.com"}`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		handler := NewGuestHandler(context.Background(),
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), mt.Client)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	mt.Run("DELETE", func(mt *mtest.T) {
		req := httptest.NewRequest(http.MethodDelete, "/submit", http.NoBody)
		rr := httptest.NewRecorder()
		handler := NewGuestHandler(context.Background(),
			slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), mt.Client)
		handler.ServeHTTP(rr, req)
		assert.Equal(t, rr.Code, http.StatusMethodNotAllowed)
		assert.Contains(t, rr.Body.String(), "invalid method")
	})
}
