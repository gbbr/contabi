package app

import (
	"net/http"

	"github.com/gbbr/contabi/tmpl"
)

func loginPage(w http.ResponseWriter, r *http.Request) error {
	return tmpl.ExecuteTemplate(w, "tmpl/files/login.html", nil)
}
