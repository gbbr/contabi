package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gbbr/contabi/tmpl"
)

func homePage(w http.ResponseWriter, r *http.Request) error {
	return tmpl.ExecuteTemplate(w, "tmpl/files/home.html", struct{ Name string }{"James"})
}

type handlerWithError struct {
	fn func(http.ResponseWriter, *http.Request) error
}

func (h handlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.fn(w, r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
}

// Serve starts the HTTP server and listens for incoming page requests.
func Serve() {
	mux := http.NewServeMux()
	mux.Handle("/", handlerWithError{homePage})

	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./assets/js"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img"))))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
