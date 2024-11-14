package main

import (
	"log"
	"net/http"
)

// serveStaticFiles serves static files such as CSS, JavaScript, and images
func serveStaticFiles() {
	// Serve files from the "static" directory at the "/static/" URL path
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

// handler serves HTML pages based on the URL path
func handler(w http.ResponseWriter, r *http.Request) {
	// Default page is "index.html" if no specific page is requested
	page := "index.html"
	switch r.URL.Path {
	case "/about":
		page = "about.html"
	case "/contact":
		page = "contact.html"
	//case "/projects":
	//	page = "projects.html"
	case "/services":
		page = "services.html"
	case "/resume":
		page = "resume.html"
	}

	// Serve the appropriate HTML file
	http.ServeFile(w, r, page)
}

func main() {
	// Serve static files (CSS, images, etc.)
	serveStaticFiles()

	// Handle HTML page requests
	http.HandleFunc("/", handler)

	// Start server on port 8080
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
