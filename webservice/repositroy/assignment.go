package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
)

// AssignmentRepository struct
type AssignmentRepository struct {
	db *sql.DB
}

// NewAssignmentRepository func
func NewAssignmentRepository(db *sql.DB) *AssignmentRepository {
	return &AssignmentRepository{
		db: db,
	}
}

// Fetch func
func (repo *AssignmentRepository) Fetch(id int) (*model.Assignment, error) {
	result := model.Assignment{}

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, reviewers FROM assignments WHERE id = ?"

	rows, err := repo.db.Query(query, id)
	if err != nil {
		return &result, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.Name, &result.Description, &result.Created, &result.Publish,
			&result.Deadline, &result.CourseID, &result.SubmissionID, &result.ReviewID, &result.Reviewers)
		if err != nil {
			return &result, err
		}
	}

	return &result, err
}

// FetchAll func
func (repo *AssignmentRepository) FetchAll() ([]*model.Assignment, error) {
	result := make([]*model.Assignment, 0)

	query := "SELECT id, name, description, created, publish, deadline, course_id, submission_id, review_id, reviewers FROM assignments"

	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.Assignment{}

		err = rows.Scan(&temp.ID, &temp.Name, &temp.Description, &temp.Created, &temp.Publish,
			&temp.Deadline, &temp.CourseID, &temp.SubmissionID, &temp.ReviewID, &temp.Reviewers)
		if err != nil {
			return result, err
		}

		result = append(result, &temp)
	}

	return result, err
}

// Insert func
func (repo *AssignmentRepository) Insert(assignment model.Assignment) (int, error) {
	var id int64

	query := "INSERT INTO assignments "

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	rows, err := tx.Exec(query)
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	id, err = rows.LastInsertId()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return int(id), err
	}

	return int(id), err
}