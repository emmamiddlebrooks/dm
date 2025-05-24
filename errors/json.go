package errors

import (
	"encoding/json"
	"net/http"
)

// JsonError a json error based on RFC 9457
type JsonError struct {
	Type     string `json:"type,omitempty"`
	Status   int    `json:"status,omitempty"`
	Title    string `json:"title,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

// CreateJsonError create a Struct with minimally required fields
func CreateJsonError(title, detail string, status int) JsonError {
	return JsonError{
		Title:  title,
		Status: status,
		Detail: detail,
	}
}

func WriteJsonError(title, detail string, status int, rw http.ResponseWriter) {
	rw.Header().Set("Content-Type", "application/problem+json")
	rw.WriteHeader(status)
	encoder := json.NewEncoder(rw)
	err := encoder.Encode(CreateJsonError(title, detail, status))
	if err != nil {
		http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
