package app

import (
	"fmt"
	"log"
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gbbr/contabi/app/tmpl"
	"github.com/gbbr/contabi/app/ui"
	"github.com/gorilla/mux"
)

type withError func(w http.ResponseWriter, r *http.Request) error

func (we withError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := we(w, r); err != nil {
		fmt.Fprintf(w, "ERROR: %+v", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) error {
	return tmpl.Execute(w, struct {
		A int
		B string
	}{1, "QWE"})
}

// Serve starts the HTTP server and listens for incoming page requests.
func Serve() {
	r := mux.NewRouter()

	// routes
	r.Handle("/", withError(home))

	// resources
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(&assetfs.AssetFS{
		Asset:     ui.Asset,
		AssetDir:  ui.AssetDir,
		AssetInfo: ui.AssetInfo,
		Prefix:    "",
	})))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
