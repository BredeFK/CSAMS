package route

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/controller"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

// Load ...
func Load() http.Handler {
	return routes()
}

// LoadHTTPS ...
func LoadHTTPS() http.Handler {
	return routes()
}

// redirectToHTTPS ....
func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, fmt.Sprintf("https://%s", r.Host), http.StatusMovedPermanently)
}

// routes setups all routes
func routes() http.Handler {
	// Instantiate mux-router
	router := mux.NewRouter().StrictSlash(true)

	// Index-page Handlers
	router.HandleFunc("/", controller.IndexGET).Methods("GET")
	router.HandleFunc("/", controller.JoinCoursePOST).Methods("POST")

	// Course-page Handlers
	router.HandleFunc("/course", controller.CourseGET).Methods("GET")
	router.HandleFunc("/course/list", controller.CourseListGET).Methods("GET")

	// Assignment-page Handlers
	router.HandleFunc("/assignment", controller.AssignmentGET).Methods("GET")
	router.HandleFunc("/assignment/peer", controller.AssignmentPeerGET).Methods("GET")
	router.HandleFunc("/assignment/auto", controller.AssignmentAutoGET).Methods("GET")

	// User-page Handlers
	router.HandleFunc("/user", controller.UserGET).Methods("GET")
	router.HandleFunc("/user/update", controller.UserUpdatePOST).Methods("POST")

	// Admin-page Handlers
	router.HandleFunc("/admin", controller.AdminGET).Methods("GET")
	router.HandleFunc("/admin/course", controller.AdminCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/create", controller.AdminCreateCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/create", controller.AdminCreateCoursePOST).Methods("POST")
	router.HandleFunc("/admin/course/update/{id}", controller.AdminUpdateCourseGET).Methods("GET")
	router.HandleFunc("/admin/course/update/{id}", controller.AdminUpdateCoursePOST).Methods("POST")
	router.HandleFunc("/admin/assignment", controller.AdminAssignmentGET).Methods("GET")
	router.HandleFunc("/admin/assignment/create", controller.AdminAssignmentCreateGET).Methods("GET")
	router.HandleFunc("/admin/assignment/create", controller.AdminAssignmentCreatePOST).Methods("POST")

	router.HandleFunc("/admin/submission", controller.AdminSubmissionGET).Methods("GET")
	router.HandleFunc("/admin/submission/create", controller.AdminSubmissionCreateGET).Methods("GET")
	router.HandleFunc("/admin/submission/create", controller.AdminSubmissionCreatePOST).Methods("POST")

	// Login/Register Handlers
	router.HandleFunc("/login", controller.LoginGET).Methods("GET")
	router.HandleFunc("/login", controller.LoginPOST).Methods("POST")
	router.HandleFunc("/register", controller.RegisterGET).Methods("GET")
	router.HandleFunc("/register", controller.RegisterPOST).Methods("POST")
	router.HandleFunc("/logout", controller.LogoutGET).Methods("GET")

	// Set path prefix for the static-folder
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// 404 Error Handler
	router.NotFoundHandler = http.HandlerFunc(controller.NotFoundHandler)
	// 405 Error Handler
	router.MethodNotAllowedHandler = http.HandlerFunc(controller.MethodNotAllowedHandler)

	return handlers.CombinedLoggingHandler(os.Stdout, router)
}
