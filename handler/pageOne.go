package handlers

import (
	"my-go-app/view/pages"
	"net/http"
)

func HandlePageOne(w http.ResponseWriter, r *http.Request) error {
	//w.Write([]byte("Hello, World! jbdjhfj"))
	//http.ServeFile(w, r, "./public/index.html")
	return pages.PageOne().Render(r.Context(), w)
}
