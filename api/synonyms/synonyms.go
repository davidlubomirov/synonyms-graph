package synonyms

import (
	"encoding/json"
	"io"
	"net/http"

	"spreadTask/api/generics"
)

type store interface {
	Get(string) []string
	Set(string, string)
}

type synonymsHandler struct {
	dataStore store
}

func NewSynonymsHandler(dataStore store) *synonymsHandler {
	return &synonymsHandler{
		dataStore,
	}
}

func (h *synonymsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	generics.SetCORS(w)

	switch {
	case r.Method == http.MethodGet:

		h.Get(w, r)

		return
	case r.Method == http.MethodPost:

		h.Set(w, r)

		return
	case r.Method == http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)

		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
}

// GET localhost:8080/synonyms?word=begin
func (h *synonymsHandler) Get(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")

	if word == "" {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	// TODO search synonyms by the the "word" query parameter
	o, _ := json.Marshal(
		&output{
			Word:     word,
			Synonyms: h.dataStore.Get(word),
		},
	)

	w.WriteHeader(http.StatusOK)
	w.Write(o)
}

// POST localhost:8080/synonyms
func (h *synonymsHandler) Set(w http.ResponseWriter, r *http.Request) {
	i := &input{}
	err := json.NewDecoder(r.Body).Decode(i)
	switch {
	case err == io.EOF:
		// empty body
		w.WriteHeader(http.StatusBadRequest)
		return
	case err != nil:
		// other error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if !i.IsValid() {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO insert the input "i" into data structure
	// ...
	h.dataStore.Set(i.Word, i.Synonym)

	w.WriteHeader(http.StatusOK)
}
