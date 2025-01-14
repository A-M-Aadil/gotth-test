package main

import (
	handlers "my-go-app/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Hello handler that returns a simple message
// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello,..    dd...")
// }

// Serve the HTML page from the public folder
// func serveHTML(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./public/index.html")
// }

// func handleLandingPage(w http.ResponseWriter, r *http.Request) {
// 	http.ServeFile(w, r, "./public/index.html")
// }

func main() {
	router := chi.NewMux()

	//router.Get("/", handleLandingPage)
	router.Get("/", handlers.HandleLandingPage)
	// Serve static HTML file (index.html)
	// http.HandleFunc("/", serveHTML) // This serves the index.html file from the public folder

	// // You can still have other route handlers if needed
	// http.HandleFunc("/hello", helloHandler) // Example of another handler

	// // Start the server on port 8080
	// fmt.Println("Server is running on port 8080 ...")
	// // Change localhost to 0.0.0.0 to allow external access
	http.ListenAndServe("0.0.0.0:8080", router)
}
