/*
The core package is used to store reusable components for any plugins or applications.
*/
package core

import (
	"html/template"
	"net/http"
	_ "net/http/pprof" // go to /debug/pprof to see debug info
)

func init() {
	http.HandleFunc("/", indexOrNotFound)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))
}

type TemplateFunc func(r *http.Request) (templatePaths []string, data interface{})

func RenderTemplate(f TemplateFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		templatePaths, data := f(r)
		templateOrPanic(data, w, r, templatePaths...)
	}
}

func indexOrNotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		templateOrPanic(r, w, r, "template/404.html")
	} else {
		templateOrPanic(nil, w, r, "template/index.html")
	}
}

func templateOrPanic(data interface{}, w http.ResponseWriter, r *http.Request, files ...string) {
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	err = t.Execute(w, r)
	if err != nil {
		panic(err)
	}
}
