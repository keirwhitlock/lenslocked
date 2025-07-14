package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my website!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to the FAQ!</h1>")
}

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
	r := chi.NewRouter()

	// chi middleware
	// https://go-chi.io/#/pages/middleware?id=logger
	r.Use(middleware.Logger)

	// chi routes
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/teapot", teapotHandler)

	// https://go-chi.io/#/pages/routing?id=routing-patterns-amp-url-parameters
	r.Get("/galleries/{id}", getGallery)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port :3000")
	http.ListenAndServe(":3000", r) // passing in nil as a Handler, uses the DefaultServeMux
}
