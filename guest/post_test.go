package guest

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPostHandler(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Run("success", func(mt *mtest.T) {
		guestString := `{"first_name":"John","last_name":"Doe","email":"johndoe@test.com"}`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		handlePost(rr, req, mt.Client)
		assert.Equal(t, rr.Code, http.StatusCreated)
	})
	mt.Run("fail decode", func(mt *mtest.T) {
		guestString := `this is not json`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		handlePost(rr, req, mt.Client)
		assert.Equal(t, rr.Code, http.StatusBadRequest)
		assert.Contains(t, rr.Body.String(), "unable to decode json")
	})
	mt.Run("fail insert", func(mt *mtest.T) {
		guestString := `{"first_name":"John","last_name":"Doe","email":"johndoe@test.com"}`
		req := httptest.NewRequest(http.MethodPost, "/submit", bytes.NewBufferString(guestString))
		rr := httptest.NewRecorder()
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "test error"}))
		handlePost(rr, req, mt.Client)
		assert.Equal(t, rr.Code, http.StatusInternalServerError)
		assert.Contains(t, rr.Body.String(), "unable to save guest info")
	})
}
