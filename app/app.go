package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

type withError struct {
	handler func(w http.ResponseWriter, r *http.Request) error
}

func (we withError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := we.handler(w, r); err != nil {
		fmt.Fprintf(w, "ERROR: %+v", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) error {
	return json.NewEncoder(w).Encode(r.Header)
}

// Serve starts the HTTP server and listens for incoming page requests.
func Serve() {
	mux := http.NewServeMux()

	// routes
	mux.Handle("/test", withError{home})

	// resources
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "",
	})))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
