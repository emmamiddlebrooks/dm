package errors

//lint:file-ignore SA5001 Ignored for test cases
import (
	"net/http/httptest"
	"testing"
)

func TestErrors_CreateJsonError(t *testing.T) {
	jsonError := CreateJsonError("Title", "Detail", 200)
	if jsonError.Title != "Title" {
		t.Errorf("unexpected title, want: %s, got: %s", "Title", jsonError.Title)
	}
	if jsonError.Detail != "Detail" {
		t.Errorf("unexpected detail, want: %s, got: %s", "Detail", jsonError.Detail)
	}
	if jsonError.Status != 200 {
		t.Errorf("unexpected status, want: %d, got: %d", 200, jsonError.Status)
	}
}

func TestErrors_WriteJsonError(t *testing.T) {
	rr := httptest.NewRecorder()
	WriteJsonError("Title", "Detail", 200, rr)
	if rr.Header().Get("Content-Type") != "application/problem+json" {
		t.Errorf("unexpected content-type, want: %s, got: %s", "application/problem+json", rr.Header().Get("Content-Type"))
	}
	expected := `{"status":200,"title":"Title","detail":"Detail"}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body, want: %s, got: %s.", expected, rr.Body.String())
	}
}
