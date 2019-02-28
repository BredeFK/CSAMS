package controller

import (
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//AssignmentGET serves assignment page to users
func AssignmentGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

//AssignmentAutoGET serves the auto validation page to user
func AssignmentAutoGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

//AssignmentPeerGET serves the peer review page to users
func AssignmentPeerGET(w http.ResponseWriter, r *http.Request) {

	//check if request has a id
	if r.Method == http.MethodGet {

		id := r.FormValue("id")
		class := r.FormValue("class")

		if id == "" || class == "" {
			//redirect to error pageinfo
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "ID is %s\nClass is %s\n", id, class)
	}

	//check that user is logged in

	//check that user is a participant in the class

	//get assignment info from database

	//parse info with template
}

// AssignmentUploadGET serves the upload page
func AssignmentUploadGET(w http.ResponseWriter, r *http.Request) {

	// Check for ID in url and give error if not
	id := r.FormValue("id")
	if id == "" {
		log.Println("Error: id can't be empty! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Convert id from string to int
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get assignment and log possible error
	assignmentRepo := model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// Get form and log possible error
	formRepo := model.FormRepository{}
	form, err := formRepo.GetFromAssignmentID(assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get course and log possible error
	course, err := model.GetCourseCodeAndName(assignment.CourseID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// TODO display answer if already uploaded
	answers, err := model.GetUserAnswers(session.GetUserFromSession(r).ID, assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	description := github_flavored_markdown.Markdown([]byte(assignment.Description))

	// Set values
	v := view.New(r)
	v.Vars["Course"] = course
	v.Vars["Assignment"] = assignment
	v.Vars["Description"] = template.HTML(description)
	v.Vars["Fields"] = form.Fields
	v.Vars["Answers"] = answers
	v.Name = "assignment/upload"
	v.Render(w)

}

// AssignmentUploadPOST servers the
func AssignmentUploadPOST(w http.ResponseWriter, r *http.Request) {

	// Check for ID in url and give error if not
	id := r.FormValue("id")
	if id == "" {
		log.Println("Error: id can't be empty! (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Convert id from string to int
	assignmentID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get assignment and log possible error
	assignmentRepo := model.AssignmentRepository{}
	assignment, err := assignmentRepo.GetSingle(assignmentID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Give error if assignment doesn't exists
	if assignment.Name == "" {
		log.Println("Error: assignment with id '" + id + "' doesn't exist! (assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	formRepo := model.FormRepository{}
	form, err := formRepo.GetFromAssignmentID(assignment.ID)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignment.SubmissionID = 1 // TODO Brede : actually get submissionID

	// Start to fill out user Submission struct
	userSub := model.UserSubmission{
		UserID:       session.GetUserFromSession(r).ID,
		SubmissionID: assignment.SubmissionID,
		AssignmentID: assignment.ID,
	}

	// Check that every form is filled an give error if not
	for _, field := range form.Fields {

		// Check if they are empty and give error if they are
		if r.FormValue(field.Name) == "" {
			log.Println("Error: assignment with form name '" + field.Name + "' can not be empty! (assignment.go)")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Add form values to the struct for user submissions
		userSub.Answers = append(userSub.Answers, model.Answer{
			Type:  field.Type,
			Value: r.FormValue(field.Name),
		})
	}

	err = model.InsertUserSubmission(userSub)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Serve front-end again
	AssignmentUploadGET(w, r)
}
