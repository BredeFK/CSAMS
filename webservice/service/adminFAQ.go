package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// FAQService struct
type FAQService struct {
	faqRepo *repository.FAQRepository
}

// NewFAQService func
func NewFAQService(db *sql.DB) *FAQService {
	return &FAQService{
		faqRepo: repository.NewFAQRepository(db),
	}
}

// TODO brede : why this? http: panic serving [::1]:8100: runtime error: invalid memory address or nil pointer dereference
// Fetch fetches the one faq in db
func (s *FAQService) Fetch() (*model.Faq, error) {
	return s.faqRepo.Fetch()
}

// Update updates the one faq in db
func (s *FAQService) Update(questions string) error {
	return s.faqRepo.Update(questions)
}
