package tmpl

import (
	"fmt"
	"io"
	"log"
	"text/template"
)

// templateFiles holds a list of file groups to be parsed. The first file in
// each group is the name of the resulting template for that group.
var templateFiles = [][]string{
	{"tmpl/files/home.html"},
	{"tmpl/files/login.html"},
}

var templates map[string]*template.Template

func init() {
	templates = make(map[string]*template.Template)
	for _, t := range templateFiles {
		if _, ok := templates[t[0]]; ok {
			log.Fatalf("duplicate template with name %s", t[0])
		}
		tpl, err := template.ParseFiles(t...)
		if err != nil {
			log.Fatalf("%v", err)
		}
		templates[t[0]] = tpl
	}
}

// ExecuteTemplate executes the template name by passing it the given data and
// writes the response to w.
func ExecuteTemplate(w io.Writer, name string, data interface{}) error {
	t, ok := templates[name]
	if !ok {
		return fmt.Errorf("template '%s' not found", name)
	}
	return t.ExecuteTemplate(w, "MAIN", data)
}
