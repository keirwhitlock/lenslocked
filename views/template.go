package views

import (
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("execute template err: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
	}
	return
}

func ParseFS(fs fs.FS, pattern ...string) (Template, error) {

	errNoSuchTemplate := template.Error{ErrorCode: template.ErrNoSuchTemplate}

	htmlTpl, err := template.ParseFS(fs, pattern...)
	if err != nil {
		e := err
		if errors.Is(err, &errNoSuchTemplate) {
			e = fmt.Errorf("template not found: %v", err)
		} else {
			e = fmt.Errorf("parsing error occured: %v", err)
		}
		return Template{}, e

	}

	return Template{
		htmlTpl: htmlTpl,
	}, nil
}
