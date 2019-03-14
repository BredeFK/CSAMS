package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/rs/xid"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func AdminChangePassGET(w http.ResponseWriter, r *http.Request) {

	// Get form value
	vars := r.FormValue("vars")

	// Only change password if vars is not empty
	if vars != "" {
		array := strings.Split(vars, "§")

		// Not enough arguments to change password
		if len(array) != 2 {
			ErrorHandler(w, r, http.StatusInternalServerError)
			log.Println("error: not enough arguments in url!")
			return
		}

		// Get password
		pass := array[0]

		// Get id and convert to int
		id, err := strconv.Atoi(array[1])
		if err != nil {
			log.Printf(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		// Update users password
		err = model.UpdateUserPassword(id, pass)
		if err != nil {
			log.Printf(err.Error())
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

	}

	// Get courses
	courseRepo := &model.CourseRepository{}
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/changepassword/index"
	v.Vars["Courses"] = courses

	v.Render(w)
}

func AdminGetUsersPOST(w http.ResponseWriter, r *http.Request) {

	// Get form value
	formVal := r.FormValue("course_id")

	// If courseID is blank, give error
	if formVal == "" {
		log.Println("error: course_id is nil")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Convert courseID to int
	courseID, err := strconv.Atoi(formVal)
	if err != nil {
		log.Printf(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get all students from courseID
	students := model.GetUsersToCourse(courseID)
	if len(students) < 0 {
		log.Println("Error: could not get students from course! (admin_change_pass.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get courses
	courseRepo := &model.CourseRepository{}
	courses, err := courseRepo.GetAllToUserSorted(session.GetUserFromSession(r).ID)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// Get new password in 20 chars
	newPass := xid.NewWithTime(time.Now()).String()

	// source: https://www.dotnetperls.com/substring-go
	// Length is 8 chars now
	safeSubstring := string([]rune(newPass)[10:18])

	// Header OK
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/changepassword/index"
	v.Vars["Courses"] = courses         // Send the courses back that the admin is teacher in
	v.Vars["Students"] = students       // Send the students in course with courseID
	v.Vars["SelectedCourse"] = courseID // Send the selected course back to fill dropdown
	v.Vars["NewPass"] = safeSubstring   // Send new password

	v.Render(w)

}
