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
		http.Error(rw, "unable to decode json", http.StatusBadRequest)
		return
	}

	err = insertGuestData(guest, client)
	if err != nil {
		http.Error(rw, "unable to save guest info", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusCreated)
}
