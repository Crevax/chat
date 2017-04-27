package templating

import (
	"html/template"
	"net/http"
	"path/filepath"
	"sync"
)

// TemplateHandler ...
type TemplateHandler struct {
	once     sync.Once
	Filename string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templating", "html", t.Filename+".html")))
	})
	t.templ.Execute(w, nil)
}
