package guest

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		guestString := `{"first_name":"John","last_name":"Doe","email":"johndoe@test.com"}`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		handlePost(rr, req)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	t.Run("fail decode", func(t *testing.T) {
		guestString := `this is not json`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		handlePost(rr, req)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Contains(t, rr.Body.String(), "unable to decode json")
	})
}
