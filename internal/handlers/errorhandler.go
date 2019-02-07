package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// ErrorHandler handles displaying errors for the clients
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)

	temp, err := template.ParseFiles("web/layout.html", "web/navbar.html", "web/error.html")

	if err != nil {
		log.Fatal(err)
	}

	if err = temp.ExecuteTemplate(w, "layout", struct {
		PageTitle   string
		LoadFormCSS bool

		ErrorCode    int
		ErrorMessage string
	}{
		PageTitle:   "Error: " + string(status),
		LoadFormCSS: true,

		ErrorCode:    status,
		ErrorMessage: http.StatusText(status),
	}); err != nil {
		log.Fatal(err)
	}

}
