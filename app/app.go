package app

import (
	"fmt"
	"log"
	"net/http"
)

// handlerWithError is an http.Handler that manages potentially returned errors.
type handlerWithError struct {
	// fn is similar to http.HandlerFunc and will be called when this handler
	// serves HTTP requests. It allows returning an error.
	fn func(http.ResponseWriter, *http.Request) error

	// noAuth, when true, will not require authentication to proceed with this
	// request.
	noAuth bool
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

	// routes
	mux.Handle("/", handlerWithError{fn: homePage})
	mux.Handle("/login", handlerWithError{fn: loginPage, noAuth: true})

	// static resources
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./assets/js"))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./assets/img"))))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
