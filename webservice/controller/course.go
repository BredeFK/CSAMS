package controller

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//CourseGET serves class page to users
func CourseGET(w http.ResponseWriter, r *http.Request) {
	//get user
	currentUser := session.GetUserFromSession(r)

	//check if user is logged in
	if !currentUser.Authenticated {
		LoginGET(w, r)
		return
	}

	vars := mux.Vars(r)
	courseID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	//get info from db
	course, err := services.Course.Fetch(courseID)
	if err != nil {
		log.Println("course service fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//assignments, err := assignmentRepo.GetAllFromCourse(courseID)
	assignments, err := services.Assignment.FetchFromCourse(course.ID)
	if err != nil {
		log.Println("get all assignments from course", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if user is an participant of said class or a teacher
	inCourse, err := services.Course.UserInCourse(currentUser.ID, courseID)
	if err != nil {
		log.Println("course service, user in course", err)
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	if !inCourse || (!inCourse && !currentUser.Teacher) {
		log.Println("user not participant of class")
		ErrorHandler(w, r, http.StatusUnauthorized)
		return
	}

	classmates, err := services.User.FetchAllFromCourse(course.ID)
	if err != nil {
		log.Println("services, user, fetch all from course", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//all a-ok
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "course"

	v.Vars["Course"] = course
	v.Vars["User"] = currentUser
	v.Vars["Classmates"] = classmates
	v.Vars["Assignments"] = assignments

	v.Render(w)
}

//CourseListGET serves class list page to users
func CourseListGET(w http.ResponseWriter, r *http.Request) {

	//check if request has an classID
	if r.Method == http.MethodGet {

		id := r.FormValue("id")

		if id == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\n", id)
	}
	//check if user is an participant of said class or a teacher

	//get classlist from db

	//parse info to html template
	temp, err := template.ParseFiles("template/test.html")
	if err != nil {
		log.Fatal(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
	}

	temp.Execute(w, nil)
}
