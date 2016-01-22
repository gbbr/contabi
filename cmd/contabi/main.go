package main

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

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleWithError(homePage))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
