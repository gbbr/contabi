package tmpl

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
)

// t holds the index HTML template.
var t *template.Template

func init() {
	b, err := Asset("index.tmpl")
	if err != nil {
		log.Fatalf("asset: %v", err)
	}

	t = template.Must(
		template.New("index").
			Funcs(template.FuncMap{
				"json": func(v interface{}) template.JS {
					a, _ := json.Marshal(v)
					return template.JS(a)
				},
			}).
			Parse(string(b)),
	)
}

func Execute(w io.Writer, v interface{}) error {
	return t.Execute(w, v)
}
