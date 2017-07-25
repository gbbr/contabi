package app

import (
	"fmt"
	"net/http"
)

func loginPage(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(w, "%s", "ASD")
	return nil
}
