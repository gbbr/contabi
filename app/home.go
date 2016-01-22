package app

import (
	"log"
	"net/http"

	"github.com/gbbr/contabi/store"
	"github.com/gbbr/contabi/tmpl"
)

func homePage(w http.ResponseWriter, r *http.Request) error {
	userStore := store.DefaultUserStore()
	u, err := userStore.Get("a@b.com")
	log.Printf("%#v %#v", u, err)
	return tmpl.ExecuteTemplate(w, "tmpl/files/home.html", struct{ Name string }{"James"})
}
