package util_test

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/handlers"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/internal/util"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetUserFromSession(t *testing.T) {
	if err := os.Chdir("../../"); err != nil { //go out of /handlers folder
		panic(err)
	}

	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated:true,
		ID:id,
		Name:name,
		EmailStudent:email,
	}

	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(handlers.MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	user2 := util.GetUserFromSession(req)

	if user2.ID != id{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", id, user2.ID)
	}
	if user2.Name != name{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", name, user2.Name)
	}
	if user2.EmailStudent != email{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", email, user2.EmailStudent)
	}
}

func TestIsLoggedIn(t *testing.T) {

	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated:true,
		ID:id,
		Name:name,
		EmailStudent:email,
	}

	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(handlers.MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	loggedIn := util.IsLoggedIn(req)

	if !loggedIn{
		t.Errorf("Not logged in expected true, got false")
	}
}

func TestIsTeacher(t *testing.T) {
	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//get a session
	session, err := db.CookieStore.Get(req, "login-session")
	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated:true,
		Teacher:true,
		ID:id,
		Name:name,
		EmailStudent:email,
	}

	//save user to session values
	session.Values["user"] = user
	//save session changes
	err = session.Save(req, resp)
	if err != nil { //check error
		t.Error(err.Error())
	}

	http.HandlerFunc(handlers.MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	isTeacher := util.IsTeacher(req)

	if !isTeacher{
		t.Errorf("Not logged in expected true, got false")
	}
}

func TestSaveUserToSession(t *testing.T) {
	id := 1
	name := "test"
	email := "hei@gmail.no"

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	//user object we want to fill with variables needed
	var user = model.User{
		Authenticated:true,
		ID:id,
		Name:name,
		EmailStudent:email,
	}

	http.HandlerFunc(handlers.MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}

	util.SaveUserToSession(user, resp, req)

	user2 := util.GetUserFromSession(req)

	if user2.ID != id{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", id, user2.ID)
	}
	if user2.Name != name{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", name, user2.Name)
	}
	if user2.EmailStudent != email{
		t.Errorf("Returned wrong user information from session, expected %v, got %v", email, user2.EmailStudent)
	}
}

