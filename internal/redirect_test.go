package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/bad-endpoint", http.NoBody)
	rr := httptest.NewRecorder()
	Redirect(rr, req)
	assert.Equal(t, rr.Code, http.StatusMovedPermanently)
}
