package main

import (
	"encoding/json"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/schedulerservice/model"
	"log"
	"net/http"
	"time"
)

type response struct {
	Success bool `json:"success"`
	Time time.Time `json:"time"`
}

// IndexGET handles GET requests
func IndexGET(w http.ResponseWriter, r *http.Request) {
	//todo
	loc, err := time.LoadLocation("Europe/Oslo")
	if err != nil{
		log.Println("Something wrong with time location")
		return
	}

	http.Header.Add(w.Header(), "content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response{Success: true, Time:time.Now().In(loc)})
}

// IndexPOST handles POST requests
func IndexPOST(w http.ResponseWriter, r *http.Request) {

	var payload model.Payload

	//decode json request into struct
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Something went wrong decoding request" + err.Error()) //todo real logger
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//schedule Task based on type of task
	switch payload.Task {
	case "peer":
		if !model.NewTask(payload) {
			log.Println("Something went wrong decoding request") //todo real logger
			http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
			return
		}
	default:
		log.Println("Something went wrong decoding request") //todo real logger
		http.Error(w, "Something went wrong decoding request", http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(response{Success: true}); err != nil {
		log.Println("Something went wrong encoding response") //todo real logger
		http.Error(w, "Something went wrong encoding response", http.StatusInternalServerError)
	}
}

// IndexPUT handles PUT requests
func IndexPUT(w http.ResponseWriter, r *http.Request) {
	//todo
}

// IndexDELETE handles DELETE requests
func IndexDELETE(w http.ResponseWriter, r *http.Request) {
	//todo
}
