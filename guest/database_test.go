package guest

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

func TestInsertGuest(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := insertGuestData(Guest{
			FirstName: "john",
			LastName:  "doe",
			Email:     "johndoe@test.com",
		}, mt.Client)
		assert.Nil(mt, err)
	})
	mt.Run("fail", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "test error"}))
		err := insertGuestData(Guest{
			FirstName: "john",
			LastName:  "doe",
			Email:     "johndoe@test.com",
		}, mt.Client)
		assert.NotNil(mt, err)
	})
}
