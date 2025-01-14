package handlers

import "net/http"

func HandleLandingPage(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello, World! jbdjhfj"))
	http.ServeFile(w, r, "./public/index.html")
}
