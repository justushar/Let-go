package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func (app *App) RenderHTML(w http.ResponseWriter, pages string, partials string) {
	files := []string{
		filepath.Join(app.HTMLDIR, "base.tmpl"),
		filepath.Join(app.HTMLDIR, pages),
		filepath.Join(app.HTMLDIR, partials),
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, err)
	}
}
