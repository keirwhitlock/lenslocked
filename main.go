package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"lenslocked/controllers"
	"lenslocked/views"
	"net/http"
)

//func executeTemplate(w http.ResponseWriter, filepath string) {
//	tpl, err := views.Parse(filepath)
//	if err != nil {
//		log.Printf("Error parsing template: %s", err)
//		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
//		return
//	}
//
//	tpl.Execute(w, nil)
//}
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, filepath.Join("templates", "home.gohtml"))
//}
//
//func contactHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, filepath.Join("templates", "contact.gohtml"))
//}
//
//func faqHandler(w http.ResponseWriter, r *http.Request) {
//	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
//}

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

	//homeTpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	//if err != nil {
	//	panic(err)
	//}
	//
	//router.Method(http.MethodGet, "/", controllers.Static{
	//	Template: homeTpl,
	//})

	// chi middleware
	// https://go-chi.io/#/pages/middleware?id=logger
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	// chi routes
	tpl := views.Must(views.Parse("templates/home.gohtml"))
	router.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/contact.gohtml"))
	router.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/faq.gohtml"))
	router.Get("/faq", controllers.StaticHandler(tpl))

	tpl = views.Must(views.Parse("templates/login.gohtml"))
	router.Get("/login", controllers.StaticHandler(tpl))

	router.Get("/teapot", teapotHandler)

	// https://go-chi.io/#/pages/routing?id=routing-patterns-amp-url-parameters
	router.Get("/galleries/{id}", getGallery)

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "404 page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on port :3000")
	http.ListenAndServe(":3000", router) // passing in nil as a Handler, uses the DefaultServeMux
}
