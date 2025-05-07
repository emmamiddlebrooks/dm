package guest

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func handlePost(rw http.ResponseWriter, req *http.Request, client *mongo.Client) {
	var guest Guest
	err := json.NewDecoder(req.Body).Decode(&guest)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	insertGuestData(guest, client)
	rw.WriteHeader(http.StatusCreated)
}
