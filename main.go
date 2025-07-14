package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// similar to fmt.Println() but you can sat where to print
	// in this case to the ResponseWriter w.
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

//func pathHandler(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/teapot":
//		teapotHandler(w, r)
//	case "/contact":
//		contactHandler(w, r)
//	case "/faq":
//		faqHandler(w, r)
//	case "/":
//		homeHandler(w, r)
//	default:
//		w.WriteHeader(http.StatusNotFound)
//		fmt.Fprint(w, "<h1>404 Not Found</h1>")
//	}
//}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/teapot":
		teapotHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	case "/":
		homeHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>404 Not Found</h1>")
	}
}

func main() {
	//var router http.HandlerFunc
	//router = pathHandler

	//http.HandleFunc("/", homeHandler) // uses the default servemux
	//http.HandleFunc("/contact", contactHandler)
	//
	var router Router

	fmt.Println("Starting the server on port :3000")
	//http.ListenAndServe(":3000", nil) // passing in nil as a Handler, uses the DefaultServeMux
	//http.ListenAndServe(":3000", http.HandlerFunc(pathHandler)) // passing in nil as a Handler, uses the DefaultServeMux
	http.ListenAndServe(":3000", router) // passing in nil as a Handler, uses the DefaultServeMux
}
