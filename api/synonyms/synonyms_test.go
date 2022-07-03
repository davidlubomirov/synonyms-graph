package synonyms

// single test only to show usage of httptest and mock

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockDataStore struct{}

func (mock *mockDataStore) Get(word string) []string {
	return []string{}
}
func (mock *mockDataStore) Set(word, synonym string) {}

func Test_synonyms_Set_invalid_http_request(t *testing.T) {
	// path can be moved as constant
	req := httptest.NewRequest(http.MethodPost, "/synonym", nil)
	w := httptest.NewRecorder()

	mockDS := &mockDataStore{}
	mockSynonymsHandler := NewSynonymsHandler(mockDS)

	mockSynonymsHandler.Set(w, req)
	res := w.Result()
	defer res.Body.Close()

	expectedStatusCode := http.StatusBadRequest
	if res.StatusCode != expectedStatusCode {
		t.Errorf("expecting status code %d, got %d", res.StatusCode, expectedStatusCode)
	}
}
