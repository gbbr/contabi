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

func handleWithError(fn func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			// Handle error page
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %v", err)
			return
		}
	}
}

// Serve starts the HTTP server and listens for incoming page requests.
func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleWithError(homePage))

	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./assets/js"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img"))))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
