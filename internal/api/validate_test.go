package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/validate", strings.NewReader(`{"keywords": ["test"]}`))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ValidateHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"matched_words": [], "count": 0}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
