package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/gbbr/contabi/app/tmpl"
	"github.com/gorilla/mux"
)

// rootDir holds the full path of the dist/ file server. It is used by
// go-bindata in development mode.
var rootDir string

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("getwd: %v", err)
	}
	rootDir = filepath.Join(wd, "app", "ui", "dist")
}

// t holds the index HTML template.
var t *template.Template

func init() {
	b, err := tmpl.Asset("index.tmpl")
	if err != nil {
		log.Fatalf("asset: %v", err)
	}

	t = template.Must(
		template.New("index").
			Funcs(template.FuncMap{
				"json": func(v interface{}) template.JS {
					a, _ := json.Marshal(v)
					return template.JS(a)
				},
			}).
			Parse(string(b)),
	)
}

type withError struct {
	handler func(w http.ResponseWriter, r *http.Request) error
}

func (we withError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := we.handler(w, r); err != nil {
		fmt.Fprintf(w, "ERROR: %+v", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) error {
	return t.Execute(w, struct {
		A int
		B string
	}{1, "QWE"})
}

// Serve starts the HTTP server and listens for incoming page requests.
func Serve() {
	r := mux.NewRouter()

	// routes
	r.Handle("/", withError{home})

	// resources
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(&assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "",
	})))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
