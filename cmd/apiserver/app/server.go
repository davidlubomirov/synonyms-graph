package app

import (
	"fmt"
	"log"
	"net/http"

	"spreadTask/api/synonyms"
	"spreadTask/internal/synonymsdatastore"
)

const (
	defaultPortNumber = 8080
)

type apiServer struct{}

func NewApiServer() *apiServer {
	return &apiServer{}
}

func (appServer *apiServer) Run() {
	handleSystemSignals()

	synonumsDataStore := synonymsdatastore.NewSynonymsStore()
	synonymsHandler := synonyms.NewSynonymsHandler(synonumsDataStore)

	http.HandleFunc("/synonym", synonymsHandler.ServeHTTP)

	fmt.Printf("starting HTTP server on port: %d\n", defaultPortNumber)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", defaultPortNumber), nil))
}
