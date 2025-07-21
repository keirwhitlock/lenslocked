package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lenslocked/controllers"
	"lenslocked/templates"
	"lenslocked/views"
	"net/http"
)

func teapotHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	fmt.Fprint(w, "<h1>Teapot</h1></br>")
}

func getGallery(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the Gallery!</h1>")
	fmt.Fprintf(w, "<p>The id was: %s</p>", id)

}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	// chi routes
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	router.Get("/faq", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "login.gohtml"))
	router.Get("/login", controllers.StaticHandler(tpl))

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port :3000")
	http.ListenAndServe(":3000", router) // passing in nil as a Handler, uses the DefaultServeMux
}
