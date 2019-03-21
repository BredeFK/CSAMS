package service

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/repositroy"
	_ "github.com/go-sql-driver/mysql" //database driver
)

// ReviewService struct
type ReviewService struct {
	reviewRepo *repositroy.ReviewRepository
	formRepo   *repositroy.FormRepository
	fieldRepo  *repositroy.FieldRepository
}

// NewReviewService func
func NewReviewService(db *sql.DB) *ReviewService {
	return &ReviewService{
		reviewRepo: repositroy.NewReviewRepository(db),
		formRepo:   repositroy.NewFormRepository(db),
		fieldRepo:  repositroy.NewFieldRepository(db),
	}
}

// FetchAll func
func (s *ReviewService) FetchAll() ([]model.Review, error) {
	result := make([]model.Review, 0)

	reviewPtr, err := s.reviewRepo.FetchAll()
	if err != nil {
		return result, err
	}

	formsPtr, err := s.formRepo.FetchAll()
	if err != nil {
		return result, err
	}

	for _, review := range reviewPtr {
		for _, form := range formsPtr {
			if review.FormID == form.ID {
				review.Form = *form
			}
		}

		result = append(result, *review)
	}

	return result, err
}

// Insert func
func (s *ReviewService) Insert(form model.Form) (int, error) {
	return s.reviewRepo.Insert(form)
}

// Update func
func (s *ReviewService) Update(form model.Form) error {
	err := s.formRepo.Update(form.ID, &form)
	if err != nil {
		return err
	}

	err = s.fieldRepo.DeleteAll(form.ID)
	if err != nil {
		return err
	}

	for _, field := range form.Fields {
		field.FormID = form.ID

		_, err = s.fieldRepo.Insert(&field)
		if err != nil {
			return err
		}
	}

	return err
}

// Delete func
func (s *ReviewService) Delete(id int) error {
	err := s.reviewRepo.Delete(id)
	if err != nil {
		return err
	}

	err = s.fieldRepo.DeleteAll(id)
	if err != nil {
		return err
	}

	err = s.formRepo.Delete(id)
	if err != nil {
		return err
	}

	return err
}