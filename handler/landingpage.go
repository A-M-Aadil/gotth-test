package handlers

import (
	"my-go-app/view/pages"
	"net/http"
)

func HandleLandingPage(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Hello, World! jbdjhfj"))
	//http.ServeFile(w, r, "./public/index.html")
	return pages.LandingPage().Render(r.Context(), w)
}
