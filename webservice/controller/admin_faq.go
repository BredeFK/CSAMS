package controller

import (
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/service"
	"github.com/JohanAanesen/CSAMS/webservice/shared/db"
	"github.com/JohanAanesen/CSAMS/webservice/shared/session"
	"github.com/JohanAanesen/CSAMS/webservice/shared/view"
	"log"
	"net/http"
)

// AdminFaqGET handles GET-request at admin/faq/index
func AdminFaqGET(w http.ResponseWriter, r *http.Request) {

	// TODO brede : use service/repository here
	content := model.GetDateAndQuestionsFAQ() // TODO (Svein): Move this to 'settings'
	if content.Questions == "-1" { // TODO (Svein): Allow blank FAQ
		log.Println("Something went wrong with getting the faq")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/index"
	v.Vars["Updated"] = content

	v.Render(w)
}

// AdminFaqEditGET returns the edit view for the faq
func AdminFaqEditGET(w http.ResponseWriter, r *http.Request) {

	// TODO brede : use service/repository here
	content := model.GetDateAndQuestionsFAQ() // TODO (Svein): Move this to 'settings'
	if content.Questions == "-1" { // TODO (Svein): Allow blank FAQ
		log.Println("Something went wrong with getting the faq")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	v := view.New(r)
	v.Name = "admin/faq/edit"
	v.Vars["Content"] = content

	v.Render(w)
}

// AdminFaqUpdatePOST handles the edited markdown faq
func AdminFaqUpdatePOST(w http.ResponseWriter, r *http.Request) {
	// Check that the questions arrived
	updatedFAQ := r.FormValue("rawQuestions")
	if updatedFAQ == "" {
		log.Println("Form is empty!")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// TODO brede : use service/repository here
	// Check that it's possible to get the old faq from db
	content := model.GetDateAndQuestionsFAQ()
	if content.Questions == "-1" {
		log.Println("Something went wrong with getting the faq")
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Check that it's changes to the new faq
	if content.Questions == updatedFAQ {
		log.Println("Old and new faq can not be equal!")
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	// TODO brede : use service/repository here
	// Check that it went okay to add new faq to db
	err := model.UpdateFAQ(updatedFAQ)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	// Get user for logging purposes
	currentUser := session.GetUserFromSession(r)

	// Services
	services := service.NewServices(db.GetDB())

	// Log update faq to db
	err = services.Logs.InsertUpdateFAQ(currentUser.ID, content.Questions, updatedFAQ)
	if err != nil {
		log.Println("log, update faq ", err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	//AdminFaqGET(w, r)
	http.Redirect(w, r, "/admin/faq", http.StatusFound)
}
