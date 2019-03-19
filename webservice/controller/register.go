package controller

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/session"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/view"
	"github.com/microcosm-cc/bluemonday"
	"log"
	"net/http"
)

//RegisterGET serves register page to users
func RegisterGET(w http.ResponseWriter, r *http.Request) {
	//course repo
	courseRepo := &model.CourseRepository{}

	name := r.FormValue("name")   // get form value name
	email := r.FormValue("email") // get form value email

	// Check if request has an courseID and it's not empty
	hash := r.FormValue("courseid")
	if hash != "" {

		// Check if the hash is a valid hash
		if course := courseRepo.CourseExists(hash); course.ID == -1 {
			ErrorHandler(w, r, http.StatusBadRequest)
			hash = ""
			return
		}
	}

	if session.IsLoggedIn(r) {
		IndexGET(w, r)
		return
	}

	v := view.New(r)
	v.Name = "register"
	// Send the correct link to template
	if hash == "" {
		v.Vars["Action"] = ""
	} else {
		v.Vars["Action"] = "?courseid=" + hash
	}

	v.Vars["Name"] = name
	v.Vars["Email"] = email

	v.Vars["Message"] = session.GetAndDeleteMessageFromSession(w, r)

	v.Render(w)

	//todo check if there is a class hash in request
	//if there is, add the user logging in to the class and redirect
}

//RegisterPOST validates register requests from users
func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	//XSS sanitizer
	p := bluemonday.UGCPolicy()

	user := session.GetUserFromSession(r)

	if user.Authenticated { //already logged in, no need to register
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	name := r.FormValue("name")         // get form value name
	email := r.FormValue("email")       // get form value email
	password := r.FormValue("password") // get form value password
	hash := r.FormValue("courseid")     // get from link courseID

	//check that nothing is empty and password match passwordConfirm
	if name == "" || email == "" || password == "" || password != r.FormValue("passwordConfirm") { //login credentials cannot be empty
		session.SaveMessageToSession("Passwords does not match or fields are empty!", w, r)
		RegisterGET(w, r)
		return
	}

	//Sanitize input
	name = p.Sanitize(name)
	email = p.Sanitize(email)
	password = p.Sanitize(password)

	user, err := model.RegisterUser(name, email, password) //register user in database
	if err != nil {
		log.Println(err.Error())
		session.SaveMessageToSession("Email already in use!", w, r)
		RegisterGET(w, r)
		return
	}

	//course repo
	courseRepo := &model.CourseRepository{}

	if user.ID != 0 {
		//save user to session values
		user.Authenticated = true
		session.SaveUserToSession(user, w, r)
		// Add new user to course

		if hash != "" {
			hash = p.Sanitize(hash)
			if id := courseRepo.CourseExists(hash).ID; id != -1 {
				courseRepo.AddUserToCourse(user.ID, id)
			}
		}
	}

	http.Redirect(w, r, "/", http.StatusFound) //success, redirect to homepage
}
