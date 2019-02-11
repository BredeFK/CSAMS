package main

/*
func main() {
	dbcon.InitDB(os.Getenv("SQLDB")) //env var SQLDB username:password@tcp(127.0.0.1:3306)/dbname 127.0.0.1 if run locally like xampp

	router := mux.NewRouter() // .StrictSlash(true) == (URL/login == URL/login/)

	router.HandleFunc("/", handlers.MainHandler).Methods("GET")

	router.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	router.HandleFunc("/login", handlers.LoginRequest).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterRequest).Methods("POST")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")

	router.HandleFunc("/course", handlers.CourseHandler).Methods("GET")
	router.HandleFunc("/course/list", handlers.CourseListHandler).Methods("GET")

	router.HandleFunc("/assignment", handlers.AssignmentHandler).Methods("GET")
	router.HandleFunc("/assignment/peer", handlers.AssignmentPeerHandler).Methods("GET")
	router.HandleFunc("/assignment/auto", handlers.AssignmentAutoHandler).Methods("GET")

	router.HandleFunc("/user", handlers.UserHandler).Methods("GET")
	router.HandleFunc("/user/update", handlers.UserUpdateRequest).Methods("POST")

	router.HandleFunc("/admin", handlers.AdminHandler).Methods("GET")
	router.HandleFunc("/admin/course", handlers.AdminCourseHandler).Methods("GET")
	router.HandleFunc("/admin/course/create", handlers.AdminCreateCourseHandler).Methods("GET")
	router.HandleFunc("/admin/course/create", handlers.AdminCreateCourseRequest).Methods("POST")
	router.HandleFunc("/admin/course/update/{id}", handlers.AdminUpdateCourseHandler).Methods("GET")
	router.HandleFunc("/admin/course/update/{id}", handlers.AdminUpdateCourseRequest).Methods("POST")
	router.HandleFunc("/admin/assignment", handlers.AdminAssignmentHandler).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(util.GetPort(), router)
}
*/
