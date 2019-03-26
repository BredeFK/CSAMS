package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/service"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/scheduler"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// AdminAssignmentGET handles GET-request at /admin/assignment
func AdminAssignmentGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())
	assignmentService := service.NewAssignmentService(db.GetDB())

	// Current user
	currentUser := session.GetUserFromSession(r)

	//get courses to user/teacher
	courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println("course service fetch all for user ordered", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//need custom struct to get the coursecode
	type ActiveAssignment struct {
		Assignment model.Assignment
		CourseCode string
	}

	var activeAssignments []ActiveAssignment

	for _, course := range courses { //iterate all courses
		assignments, err := assignmentService.FetchFromCourse(course.ID)
		if err != nil {
			log.Println("fetch from course", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		for _, assignment := range assignments { //go through all it's assignments again
			activeAssignments = append(activeAssignments, ActiveAssignment{Assignment: *assignment, CourseCode: course.Code})
		}

	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/index"

	v.Vars["Assignments"] = activeAssignments

	v.Render(w)
}

// AdminAssignmentCreateGET handles GET-request from /admin/assigment/create
func AdminAssignmentCreateGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())
	submissionService := service.NewSubmissionService(db.GetDB())
	reviewService := service.NewReviewService(db.GetDB())

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Fetch all submission
	submissions, err := submissionService.FetchAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch all reviews
	reviews, err := reviewService.FetchAll()
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch courses, ordered
	courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println("course service, fetch all for user ordered", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/create"

	v.Vars["Courses"] = courses

	v.Vars["Submissions"] = submissions
	v.Vars["Reviews"] = reviews

	v.Render(w)
}

// AdminAssignmentCreatePOST handles POST-request from /admin/assigment/create
func AdminAssignmentCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())
	assignmentService := service.NewAssignmentService(db.GetDB())

	// Current user
	currentUser := session.GetUserFromSession(r)

	// Declare empty assignment
	assignment := model.Assignment{}

	// Declare empty slice of strings
	var errorMessages []string

	// Get form name from request
	assignmentName := r.FormValue("name")
	// Get form description from request
	assignmentDescription := r.FormValue("description")

	// Check if name is empty
	if assignmentName == "" {
		errorMessages = append(errorMessages, "Error: Assignment Name cannot be blank.")
	}

	// Get the time.Time object from the publish string
	publish, err := util.DatetimeLocalToRFC3339(r.FormValue("publish"))
	if err != nil {
		errorMessages = append(errorMessages, "Error: Something wrong with the publish datetime.")
	}

	// Get the time.Time object from the deadline string
	deadline, err := util.DatetimeLocalToRFC3339(r.FormValue("deadline"))
	if err != nil {
		errorMessages = append(errorMessages, "Error: Something wrong with the deadline datetime.")
	}

	// Check if publish datetime is after the deadline
	if publish.After(deadline) {
		errorMessages = append(errorMessages, "Error: Deadline cannot be before Publish.")
	}

	// Check if there are any error messages
	if len(errorMessages) != 0 {
		// TODO (Svein): Keep data from the previous submit
		courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
		if err != nil {
			log.Println("course service, fetch all for user ordered", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		v := view.New(r)
		v.Name = "admin/assignment/create"

		v.Vars["Errors"] = errorMessages
		v.Vars["AssignmentName"] = assignmentName
		v.Vars["AssignmentDescription"] = assignmentDescription
		v.Vars["Courses"] = courses

		v.Render(w)
		return
	}

	// Get form values
	var val int

	// String converted into integer
	courseID, err := strconv.Atoi(r.FormValue("course_id"))
	if err != nil {
		log.Print("course_id")
		log.Println(err)
		return
	}

	if r.FormValue("submission_id") != "" {
		val, err = strconv.Atoi(r.FormValue("submission_id"))
		if err != nil {
			log.Println("submission_id")
			log.Println(err)
			return
		}
	}
	submissionID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	val = 0

	if r.FormValue("review_id") != "" {
		val, err = strconv.Atoi(r.FormValue("review_id"))
		if err != nil {
			log.Println("review_id")
			log.Println(err)
			return
		}

		// Get the time.Time object from the deadline string
		reviewDeadline, err := util.DatetimeLocalToRFC3339(r.FormValue("review_deadline"))
		if err != nil {
			errorMessages = append(errorMessages, "Error: Something wrong with the review deadline datetime.")
		}
		// Put review deadline into assignment
		assignment.ReviewDeadline = reviewDeadline
	}
	reviewID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	if r.FormValue("reviewers") != "" {
		val, err = strconv.Atoi(r.FormValue("reviewers"))
		if err != nil {
			log.Println("reviewers")
			log.Println(err)
			return
		}
	}
	reviewers := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	// Put all data into an Assignment-struct
	assignment.Name = assignmentName
	assignment.Description = assignmentDescription
	assignment.Publish = publish
	assignment.Deadline = deadline
	assignment.CourseID = courseID
	assignment.SubmissionID = submissionID
	assignment.ReviewID = reviewID
	assignment.Reviewers = reviewers

	// Insert data to database
	lastID, err := assignmentService.Insert(assignment)
	if err != nil {
		log.Println("assignment service, insert", err)
		return
	}

	// if submission ID AND Reviewers is set and valid, we can schedule the peer_review service to execute  TODO time-norwegian
	if lastID != 0 && assignment.SubmissionID.Valid && assignment.Reviewers.Valid && assignment.Deadline.After(util.GetTimeInCorrectTimeZone()) {

		sched := scheduler.Scheduler{}

		err := sched.SchedulePeerReview(int(assignment.SubmissionID.Int64),
			lastID, //assignment ID
			int(assignment.Reviewers.Int64),
			assignment.Deadline)
		if err != nil {
			log.Println(err)
			return
		}

	}

	http.Redirect(w, r, "/admin/assignment", http.StatusFound)
}

// AdminSingleAssignmentGET handles GET-request at admin/assignment/{id}
func AdminSingleAssignmentGET(w http.ResponseWriter, r *http.Request) {
	// Services
	assignmentService := service.NewAssignmentService(db.GetDB())

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignment, err := assignmentService.Fetch(id)
	if err != nil {
		log.Println("assignment service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	descriptionMD := []byte(assignment.Description)
	description := github_flavored_markdown.Markdown(descriptionMD)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/single"

	v.Vars["Assignment"] = assignment
	v.Vars["Description"] = template.HTML(description) // TODO (Svein): User template function

	v.Render(w)
}

// AdminUpdateAssignmentGET handles GET-request at /admin/assignment/update/{id}
func AdminUpdateAssignmentGET(w http.ResponseWriter, r *http.Request) {
	// Services
	courseService := service.NewCourseService(db.GetDB())
	assignmentService := service.NewAssignmentService(db.GetDB())
	submissionService := service.NewSubmissionService(db.GetDB())
	reviewService := service.NewReviewService(db.GetDB())
	submissionAnswerService := service.NewSubmissionAnswerService(db.GetDB())

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get current user
	currentUser := session.GetUserFromSession(r)

	// Fetch all submissions
	submissions, err := submissionService.FetchAll()
	if err != nil {
		log.Println("submission service, fetch all", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch assignment
	assignment, err := assignmentService.Fetch(id)
	if err != nil {
		log.Println("assignment service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get number of Students that has delivered submission with specific submission form
	submissionCount, err := submissionAnswerService.CountForAssignment(assignment.ID)
	if err != nil {
		log.Println("submission answer service, count for assignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get courses to user
	courses, err := courseService.FetchAllForUserOrdered(currentUser.ID)
	if err != nil {
		log.Println("course service, fetch all for user ordered", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch all reviews
	reviews, err := reviewService.FetchAll()
	if err != nil {
		log.Println("review service, fetch all", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/update"

	v.Vars["Assignment"] = assignment
	v.Vars["SubmissionCount"] = submissionCount
	v.Vars["Publish"] = util.GoToHTMLDatetimeLocal(assignment.Publish)
	v.Vars["Deadline"] = util.GoToHTMLDatetimeLocal(assignment.Deadline)
	if !assignment.ReviewDeadline.IsZero() {
		v.Vars["ReviewDeadline"] = util.GoToHTMLDatetimeLocal(assignment.ReviewDeadline)
	}
	v.Vars["Courses"] = courses
	v.Vars["Submissions"] = submissions
	v.Vars["Reviews"] = reviews

	v.Render(w)
}

// AdminUpdateAssignmentPOST handles POST-request at /admin/assignment/update
func AdminUpdateAssignmentPOST(w http.ResponseWriter, r *http.Request) {
	// Services
	assignmentService := service.NewAssignmentService(db.GetDB())
	submissionAnswerService := service.NewSubmissionAnswerService(db.GetDB())

	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Empty assignment
	assignment := model.Assignment{}

	// Get the time.Time object from the publish string
	publish, err := util.DatetimeLocalToRFC3339(r.FormValue("publish"))
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get the time.Time object from the deadline string
	deadline, err := util.DatetimeLocalToRFC3339(r.FormValue("deadline"))
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if publish datetime is after the deadline
	if publish.After(deadline) {
		log.Println("publish after deadline")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form values
	courseIDString := r.FormValue("course_id")
	var val int

	// String converted into integer
	courseID, err := strconv.Atoi(courseIDString)
	if err != nil {
		log.Printf("course_id: %v", err)
		return
	}

	if r.FormValue("submission_id") != "" {
		val, err = strconv.Atoi(r.FormValue("submission_id"))
		if err != nil {
			log.Println("submission_id")
			log.Println(err)
			return
		}
	}

	submissionID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	// Delete former submissions if admin changes submission form
	formerAssignment, err := assignmentService.Fetch(id)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	var formerID int64
	var newID int64
	formerID = 0
	newID = 0

	if formerAssignment.SubmissionID.Valid {
		formerID = formerAssignment.SubmissionID.Int64
	}
	if submissionID.Valid {
		newID = submissionID.Int64
	}

	// If submission id has changed, and it wasn't 'None' before, delete former submissions
	if formerID != newID && formerID != 0 {
		err = submissionAnswerService.DeleteFromAssignment(formerAssignment.ID)
		if err != nil {
			log.Println("submission answer service, delete from assignment", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
	}

	val = 0

	if r.FormValue("review_id") != "" {
		val, err = strconv.Atoi(r.FormValue("review_id"))
		if err != nil {
			log.Println("string convert review_id", err)
			return
		}

		// Get the time.Time object from the deadline string
		reviewDeadline, err := util.DatetimeLocalToRFC3339(r.FormValue("review_deadline"))
		if err != nil {
			log.Println(err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		assignment.ReviewDeadline = reviewDeadline
	}

	reviewID := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	if r.FormValue("reviewers") != "" {
		val, err = strconv.Atoi(r.FormValue("reviewers"))
		if err != nil {
			log.Println("string convert reviewers", err)
			return
		}
	}
	reviewers := sql.NullInt64{
		Int64: int64(val),
		Valid: val != 0,
	}

	assignment.ID = id
	assignment.Name = r.FormValue("name") // TODO (Svein): Sanitize
	assignment.Description = r.FormValue("description")
	assignment.Publish = publish
	assignment.Deadline = deadline
	assignment.CourseID = courseID
	assignment.SubmissionID = submissionID
	assignment.ReviewID = reviewID
	assignment.Reviewers = reviewers

	err = assignmentService.Update(assignment)
	if err != nil {
		log.Println(err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// if submission ID AND Reviewers is set and valid, we can schedule the peer_review service to execute TODO time-norwegian
	if assignment.ID != 0 && assignment.SubmissionID.Valid && assignment.Reviewers.Valid && assignment.Deadline.After(util.GetTimeInCorrectTimeZone()) {
		sched := scheduler.Scheduler{}

		if sched.SchedulerExists(int(assignment.SubmissionID.Int64), assignment.ID) {
			err := sched.UpdateSchedule(int(assignment.SubmissionID.Int64),
				assignment.ID, //assignment ID
				int(assignment.Reviewers.Int64),
				assignment.Deadline)
			if err != nil {
				log.Println(err)
				return
			}
		} else {
			err := sched.SchedulePeerReview(int(assignment.SubmissionID.Int64),
				assignment.ID, //assignment ID
				int(assignment.Reviewers.Int64),
				assignment.Deadline)
			if err != nil {
				log.Println(err)
				return
			}
		}

	}

	http.Redirect(w, r, "/admin/assignment", http.StatusFound)
}

// AdminAssignmentSubmissionsGET servers list of all users in course to admin
func AdminAssignmentSubmissionsGET(w http.ResponseWriter, r *http.Request) {
	// Services
	services := service.NewServices(db.GetDB())

	vars := mux.Vars(r)
	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println("strconv, id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("assignment service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("course service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	submissionCount, err := services.SubmissionAnswer.CountForAssignment(assignment.ID)
	if err != nil {
		log.Println("submission answer service, count for assignment", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	students, err := services.User.FetchAllStudentsFromCourse(assignment.CourseID)
	if err != nil {
		log.Println("user service, fetch all from course", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	if len(students) < 0 {
		log.Println("Error: could not get students from course! (admin_assignment.go)")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	type UserSubmissionData struct {
		User          model.User
		SubmittedTime time.Time
		Submitted     bool
		Reviewed      bool
	}

	var users []UserSubmissionData

	for _, student := range students {
		submitTime, submitted, err := services.SubmissionAnswer.FetchSubmittedTime(student.ID, assignment.ID)
		if err != nil {
			log.Println("services submission answer, fetch submitted time", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		reviews, err := services.ReviewAnswer.FetchForTarget(student.ID, assignment.ID)
		if err != nil {
			log.Println("services, review answer, fetch for target", err)
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}

		var data = UserSubmissionData{
			User:      *student,
			Submitted: false,
			Reviewed:  len(reviews) > 0,
		}

		if submitted {
			data.SubmittedTime = submitTime
			data.Submitted = true
		}

		users = append(users, data)
	}

	//Sort slice by submitted time
	sort.Slice(users, func(i, j int) bool {
		return users[i].SubmittedTime.After(users[j].SubmittedTime)
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/assignment/submissions"

	v.Vars["SubmissionCount"] = submissionCount
	v.Vars["Assignment"] = assignment
	v.Vars["Students"] = users
	v.Vars["Course"] = course

	v.Render(w)
}

/*
TODO brede : use this with iframe after alpha
// AdminAssignmentSubmissionGET servers one user submission in course to admin
func AdminAssignmentSubmissionGET(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	assignmentID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("id: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(r.FormValue("userid"))
	if err != nil {
		log.Printf("userid: %v", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	assignmentRepo := &model.AssignmentRepository{}

	assignment, err := assignmentRepo.GetSingle(int(assignmentID))
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// TODO brede : use same page as peer rews aka. out of admin/
	v := view.New(r)
	v.Name = "admin/assignment/singlesubmission"

	v.Vars["Assignment"] = assignment

	v.Render(w)

}
*/

// AdminAssignmentReviewsGET handles request to /admin/assignment/{id}/review/{id}
func AdminAssignmentReviewsGET(w http.ResponseWriter, r *http.Request) {
	// Services
	userService := service.NewUserService(db.GetDB())
	reviewAnswerService := service.NewReviewAnswerService(db.GetDB())

	// Get URL variables
	vars := mux.Vars(r)
	// Get assignment id
	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("string convert assignment id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user id
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("string convert user id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	reviews, err := reviewAnswerService.FetchForUser(userID, assignmentID)
	if err != nil {
		log.Println("review answer service, fetch for target", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user data
	user, err := userService.Fetch(userID)
	if err != nil {
		log.Println("user service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	type WeightData struct {
		IsWeighted bool
		Total      float32
		Score      float32
		Percent    float32
	}

	weights := make([]WeightData, 0)

	for _, review := range reviews {
		var totalWeight float32
		var weightScore float32

		for _, item := range review {
			totalWeight += float32(item.Weight)
			if item.Type == "checkbox" {
				if item.Answer == "on" {
					weightScore += float32(item.Weight)
				}

			} else if item.Type == "radio" {
				for k := range item.Choices {
					ans, _ := strconv.Atoi(item.Answer)
					if ans == (k + 1) {
						K := float32(k + 1)
						L := float32(len(item.Choices))
						V := float32(item.Weight) * (K / L)
						weightScore += V
					}
				}
			}
		}

		weights = append(weights, WeightData{
			IsWeighted: totalWeight > 0,
			Total:      totalWeight,
			Score:      weightScore,
			Percent:    (weightScore / totalWeight) * 100.0,
		})
	}

	totalWeights := WeightData{
		IsWeighted: false,
	}
	var sum float32
	for _, item := range weights {
		sum += item.Score
	}

	if sum > 0 {
		avg := sum / float32(len(weights))
		totalWeights.IsWeighted = true
		totalWeights.Score = avg
		totalWeights.Total = weights[0].Total
		totalWeights.Percent = (totalWeights.Score / totalWeights.Total) * 100.0
	}

	// Set header content-type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	// Set template
	v.Name = "admin/assignment/reviews"

	// View variables
	v.Vars["AssignmentID"] = assignmentID
	v.Vars["User"] = user
	v.Vars["Reviews"] = reviews
	v.Vars["Weights"] = weights
	v.Vars["TotalWeights"] = totalWeights

	// Render view
	v.Render(w)
}

// AdminAssignmentSingleSubmissionGET handles GET-request at /admin/assignment/{id}/submission/{id}
func AdminAssignmentSingleSubmissionGET(w http.ResponseWriter, r *http.Request) {
	// Services
	submissionService := service.NewSubmissionService(db.GetDB())
	submissionAnswerService := service.NewSubmissionAnswerService(db.GetDB())
	userService := service.NewUserService(db.GetDB())

	// Get URL variables
	vars := mux.Vars(r)
	// Get assignment id
	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("string convert assignment id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user id
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("string convert user id", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//user := model.GetUser(userID)
	user, err := userService.Fetch(userID)
	if err != nil {
		log.Println("user service, fetch", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	_, err = submissionService.FetchFromAssignment(assignmentID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	//answers, err := model.GetUserAnswers(user.ID, assignmentID)
	answers, err := submissionAnswerService.FetchUserAnswers(user.ID, assignmentID)
	if err != nil {
		log.Println("get user answers", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	var totalWeight float32
	var weightScore float32
	var scorePercent float32
	var isWeighted = false

	for _, item := range answers {
		totalWeight += float32(item.Weight)
		if item.Type == "checkbox" {
			if item.Answer == "on" {
				weightScore += float32(item.Weight)
			}

		} else if item.Type == "radio" {
			for k := range item.Choices {
				ans, _ := strconv.Atoi(item.Answer)
				if ans == (k + 1) {
					K := float32(k + 1)
					L := float32(len(item.Choices))
					V := float32(item.Weight) * (K / L)
					weightScore += V
				}
			}
		}
	}

	if totalWeight > 0 {
		isWeighted = true
	}

	scorePercent = (weightScore / totalWeight) * 100

	// Set header content-type and status code
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	// Set template
	v.Name = "admin/assignment/singlesubmission"

	// View variables
	v.Vars["AssignmentID"] = assignmentID
	v.Vars["User"] = user
	v.Vars["Answers"] = answers
	v.Vars["TotalWeight"] = totalWeight
	v.Vars["WeightScore"] = weightScore
	v.Vars["ScorePercent"] = scorePercent
	v.Vars["IsWeighted"] = isWeighted

	// Render view
	v.Render(w)
}

// AdminAssignmentSubmissionCreateGET func
func AdminAssignmentSubmissionCreateGET(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("strconv, atoi, assignmentID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("strconv, atoi, userID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Fetch user
	user, err := services.User.Fetch(userID)
	if err != nil {
		log.Println("services, user, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	submitted, err := services.SubmissionAnswer.HasUserSubmitted(assignment.ID, user.ID)
	if err != nil {
		log.Println("services, submission answer, has user submitted", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if submitted {
		http.Redirect(w, r,
			fmt.Sprintf("/admin/assignment/%d/submissions", assignment.ID),
			http.StatusTemporaryRedirect)
		return
	}

	// Fetch course
	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("services, course, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	answers, err := services.SubmissionAnswer.FetchUserAnswers(user.ID, assignment.ID)
	if err != nil {
		log.Println("services, submission answer, fetch user answers", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if len(answers) == 0 {
		for _, item := range submissionForm.Form.Fields {
			answers = append(answers, &model.SubmissionAnswer{
				Type:        item.Type,
				Choices:     item.Choices,
				Weight:      item.Weight,
				Label:       item.Label,
				HasComment:  item.HasComment,
				Description: item.Description,
				Name:        item.Name,
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	v.Name = "admin/assignment/submission/create"

	v.Vars["User"] = user
	v.Vars["Assignment"] = assignment
	v.Vars["Course"] = course
	v.Vars["Answers"] = answers

	v.Render(w)
}

// AdminAssignmentSubmissionCreatePOST func
func AdminAssignmentSubmissionCreatePOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()

	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("strconv, atoi, assignmentID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("strconv, atoi, userID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Fetch user
	user, err := services.User.Fetch(userID)
	if err != nil {
		log.Println("services, user, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Create empty submission answer slice
	submissionAnswers := make([]*model.SubmissionAnswer, 0)

	// Parse form from request
	err = r.ParseForm()
	if err != nil {
		log.Println("request parse form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that every form is filled an give error if not
	for _, field := range submissionForm.Form.Fields {
		// Check if they are empty and give error if they are
		if r.FormValue(field.Name) == "" && field.Type != "checkbox" && field.Type != "paragraph" && field.Type != "multi-checkbox" {
			log.Println("Error: assignment with form name '" + field.Name + "' can not be empty! (assignment.go)")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Initialize empty answer
		answer := model.SubmissionAnswer{}

		// Check if multi-checkbox
		if field.Type == "multi-checkbox" {
			val := r.Form[field.Name]
			answer.Answer = p.Sanitize(strings.Join(val, ","))
		} else {
			// Sanitize input
			answer.Answer = p.Sanitize(r.FormValue(field.Name))
		}

		// Get field type
		answer.UserID = user.ID
		answer.AssignmentID = assignment.ID
		answer.SubmissionID = int(assignment.SubmissionID.Int64)
		answer.Type = field.Type
		answer.Name = field.Name
		answer.Label = field.Label
		answer.HasComment = field.HasComment
		//answer.Description = field.Description
		// Check if the field has comment enabled
		if field.HasComment {
			// Get comment content, sanitized
			answer.Comment.String = p.Sanitize(r.FormValue(field.Name + "_comment"))
			answer.Comment.Valid = answer.Comment.String != ""
		}

		// If delivered, only change the value
		submissionAnswers = append(submissionAnswers, &answer)
	}

	err = services.SubmissionAnswer.Insert(submissionAnswers)
	if err != nil {
		log.Println("services, submission answer, insert", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf(
			"/admin/assignment/%d/submissions",
			assignment.ID),
		http.StatusFound)
}

// AdminAssignmentSubmissionUpdateGET func
func AdminAssignmentSubmissionUpdateGET(w http.ResponseWriter, r *http.Request) {
	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("strconv, atoi, assignmentID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("strconv, atoi, userID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Fetch user
	user, err := services.User.Fetch(userID)
	if err != nil {
		log.Println("services, user, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	submitted, err := services.SubmissionAnswer.HasUserSubmitted(assignment.ID, user.ID)
	if err != nil {
		log.Println("services, submission answer, has user submitted", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check if user has not submitted
	if !submitted {
		http.Redirect(w, r,
			fmt.Sprintf("/admin/assignment/%d/submissions", assignment.ID),
			http.StatusTemporaryRedirect)
		return
	}

	// Fetch course
	course, err := services.Course.Fetch(assignment.CourseID)
	if err != nil {
		log.Println("services, course, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get answers to user if he has delivered
	answers, err := services.SubmissionAnswer.FetchUserAnswers(user.ID, assignment.ID)
	if err != nil {
		log.Println("services, submission answer, fetch user answers", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if len(answers) == 0 {
		for _, item := range submissionForm.Form.Fields {
			answers = append(answers, &model.SubmissionAnswer{
				Type:        item.Type,
				Choices:     item.Choices,
				Weight:      item.Weight,
				Label:       item.Label,
				HasComment:  item.HasComment,
				Description: item.Description,
				Name:        item.Name,
			})
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	// Create view
	v := view.New(r)
	v.Name = "admin/assignment/submission/update"

	v.Vars["User"] = user
	v.Vars["Assignment"] = assignment
	v.Vars["Course"] = course
	v.Vars["Answers"] = answers

	v.Render(w)
}

// AdminAssignmentSubmissionUpdatePOST func
func AdminAssignmentSubmissionUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Sanitizer
	p := bluemonday.UGCPolicy()

	// Get URL variables
	vars := mux.Vars(r)

	assignmentID, err := strconv.Atoi(vars["assignmentID"])
	if err != nil {
		log.Println("strconv, atoi, assignmentID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		log.Println("strconv, atoi, userID", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	// Fetch user
	user, err := services.User.Fetch(userID)
	if err != nil {
		log.Println("services, user, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Fetch assignment
	assignment, err := services.Assignment.Fetch(assignmentID)
	if err != nil {
		log.Println("services, assignment, fetch", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get form and log possible error
	submissionForm, err := services.Submission.FetchFromAssignment(assignment.ID)
	if err != nil {
		log.Println("get submission form from assignment id", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Create empty submission answer slice
	submissionAnswers, err := services.SubmissionAnswer.FetchUserAnswers(user.ID, assignment.ID)
	if err != nil {
		log.Println("services, submission answers, fetch user answers", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Parse form from request
	err = r.ParseForm()
	if err != nil {
		log.Println("request parse form", err)
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that every form is filled an give error if not
	for index, field := range submissionForm.Form.Fields {
		// Check if they are empty and give error if they are
		if r.FormValue(field.Name) == "" && field.Type != "checkbox" && field.Type != "paragraph" && field.Type != "multi-checkbox" {
			log.Println("Error: assignment with form name '" + field.Name + "' can not be empty! (assignment.go)")
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}

		// Initialize empty answer
		answer := model.SubmissionAnswer{}
		// Check if the field has comment enabled
		if field.HasComment {
			// Get comment content, sanitized
			answer.Comment.String = p.Sanitize(r.FormValue(field.Name + "_comment"))
			answer.Comment.Valid = answer.Comment.String != ""
		}

		// Check if multi-checkbox
		if field.Type == "multi-checkbox" {
			val := r.Form[field.Name]
			answer.Answer = p.Sanitize(strings.Join(val, ","))
		} else {
			// Sanitize input
			answer.Answer = p.Sanitize(r.FormValue(field.Name))
		}

		// Get field type
		answer.Type = field.Type
		answer.Name = field.Name
		answer.Label = field.Label
		answer.Description = field.Description
		answer.HasComment = field.HasComment

		// If delivered, only change the value
		submissionAnswers[index].Answer = answer.Answer
		submissionAnswers[index].Comment = answer.Comment
		// Set name & label
		submissionAnswers[index].Name = field.Name
		submissionAnswers[index].Label = field.Label
		submissionAnswers[index].Description = field.Description
		submissionAnswers[index].HasComment = field.HasComment
	}

	// Update user, assignment & submission id for all answers
	for _, item := range submissionAnswers {
		item.UserID = user.ID
		item.AssignmentID = assignment.ID
		item.SubmissionID = int(assignment.SubmissionID.Int64)
	}

	err = services.SubmissionAnswer.Update(submissionAnswers)
	if err != nil {
		log.Println("services, submission answer, insert", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r,
		fmt.Sprintf(
			"/admin/assignment/%d/submissions",
			assignment.ID),
		http.StatusFound)
}

// AdminAssignmentSubmissionDELETE func
func AdminAssignmentSubmissionDELETE(w http.ResponseWriter, r *http.Request) {
	// Respond struct
	respond := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{}

	body := struct {
		AssignmentID int `json:"assignment_id"`
		UserID       int `json:"user_id"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		respond.Code = http.StatusBadRequest
		respond.Message = "Could not decode request body"

		w.WriteHeader(respond.Code)
		err = json.NewEncoder(w).Encode(respond)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
			return
		}
		return
	}

	// Services
	services := service.NewServices(db.GetDB())

	err = services.SubmissionAnswer.Delete(body.AssignmentID, body.UserID)
	if err != nil {
		respond.Code = http.StatusInternalServerError
		respond.Message = "Could not delete submission"

		w.WriteHeader(respond.Code)
		err = json.NewEncoder(w).Encode(respond)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Something went wrong."))
			return
		}
		return
	}

	respond.Code = http.StatusOK
	respond.Message = "Submission deleted successfully"

	w.WriteHeader(respond.Code)
	err = json.NewEncoder(w).Encode(respond)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong."))
		return
	}
}