package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func init() {
	if err := os.Chdir("../"); err != nil { //go out of /handlers folder
		panic(err)
	}
}

func TestMainHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(MainHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestLoginHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(LoginHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestLogoutHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(LogoutHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestClassHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class?id=asdfcvgbhnjk", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(ClassHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestClassListHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/class/list?id=adsikjuh", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(ClassListHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestUserHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/user", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(UserHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAdminHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AdminHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentAutoHandlerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/auto?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentAutoHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestAssignmentPeerHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/assignment/peer?id=ihadls&class=asdbjlid", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentPeerHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusOK, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}

func TestErrorHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/class", nil) //class with no id should give 403
	if err != nil {
		t.Fatal(err.Error())
	}

	resp := httptest.NewRecorder()

	http.HandlerFunc(AssignmentPeerHandler).ServeHTTP(resp, req)

	status := resp.Code

	if status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code, expected %v, got %v", http.StatusBadRequest, status)
	}

	body := resp.Body

	if body.Len() <= 0 {
		t.Errorf("Response body error, expected greater than 0, got %d", body.Len())
	}
}