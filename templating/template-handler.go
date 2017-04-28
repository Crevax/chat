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
	Filepath []string
	templ    *template.Template
}

func (t *TemplateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join(t.Filepath...)))
	})
	t.templ.Execute(w, nil)
}
