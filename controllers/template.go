package controllers

import "net/http"

type Template interface {
	Execute(w http.ResponseWriter, data interface{}) // What args does this take?
}
