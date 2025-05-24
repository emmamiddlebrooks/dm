package guest

import (
	"encoding/json"
	"fmt"
	"gfi/errors"
	"net/http"
	"strings"
)

const sheetsUrl string = "https://script.google.com/macros/s/AKfycbwj41ElKQF0ss79_MGQLgafLhHvLSeFGvgn71W7FtS_XttEytu4QkUSThHtQczcHths/exec"

func handlePost(rw http.ResponseWriter, req *http.Request) {
	var guest Request
	err := json.NewDecoder(req.Body).Decode(&guest)
	if err != nil {
		errors.WriteJsonError(http.StatusText(http.StatusBadRequest),
			"unable to decode json", http.StatusBadRequest, rw)
		return
	}
	if guest.WifiPassword != "BuyABeer" {
		errors.WriteJsonError(http.StatusText(http.StatusBadRequest),
			"invalid password", http.StatusBadRequest, rw)
		return
	}
	err = insertGuestData(guest)
	if err != nil {
		errors.WriteJsonError(http.StatusText(http.StatusInternalServerError),
			"unable to save guest info", http.StatusInternalServerError, rw)
		return
	}
	rw.WriteHeader(http.StatusCreated)
	redirectUrl := "http://10.1.10.165:8880/guest/s/default/authorize?id=" + guest.ClientID
	err = json.NewEncoder(rw).Encode(map[string]string{"redirect_url": redirectUrl})
	if err != nil {
		errors.WriteJsonError(http.StatusText(http.StatusInternalServerError),
			"unable to send redirect url", http.StatusInternalServerError, rw)
	}
}

func insertGuestData(guest Request) error {
	client := &http.Client{}
	data, err := json.Marshal(guest)
	if err != nil {
		return fmt.Errorf("error marshaling json: %v", err)
	}
	req, err := http.NewRequest("POST", sheetsUrl, strings.NewReader(string(data)))
	if err != nil {
		return fmt.Errorf("error sending request to signup sheet: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request to signup sheet: %v", err)
	}
	return nil
}
