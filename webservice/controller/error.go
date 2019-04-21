package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"net/http"
)

// ErrorHandler handles displaying errors for the clients
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	v := view.New(r)
	v.Name = "error"

	v.Vars["ErrorCode"] = status
	v.Vars["ErrorMessage"] = http.StatusText(status)

	if status == 99999{
		v.Vars["ErrorCode"] = "Review Error"
		v.Vars["ErrorMessage"] = "You have reviewed every submission on this assignment!"
	}

	v.Render(w)
}

// NotFoundHandler handles 404 errors
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, http.StatusNotFound)
}

// MethodNotAllowedHandler handles 405 errors
func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, r, http.StatusMethodNotAllowed)
}
