package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gbbr/contabi/svc"
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

func authed(r *http.Request) bool {
	ok, err := svc.Users.IsValidRequest(r)
	if err != nil {
		return false
	}
	return ok
}

func (h handlerWithError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.noAuth && !authed(r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
	if err := h.fn(w, r); err != nil {
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
