package app

import (
	"log"
	"net/http"

	"github.com/gbbr/contabi/tmpl"
)

func loginPage(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			log.Printf("%v", err)
		}
		usr, pass := r.Form.Get("user"), r.Form.Get("pass")
		log.Printf("%s/%s", usr, pass)
	}
	return tmpl.ExecuteTemplate(w, "tmpl/files/login.html", nil)
}
