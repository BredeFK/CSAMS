package repositroy

import (
	"database/sql"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/model"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"strings"
)

// ReviewAnswerRepository struct
type ReviewAnswerRepository struct {
	db *sql.DB
}

// NewReviewAnswerRepository func
func NewReviewAnswerRepository(db *sql.DB) *ReviewAnswerRepository {
	return &ReviewAnswerRepository{
		db: db,
	}
}

func fetchMany(repo *ReviewAnswerRepository, query string, args ...interface{}) ([]*model.ReviewAnswer, error) {
	result := make([]*model.ReviewAnswer, 0)

	rows, err := repo.db.Query(query, args...)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		temp := model.ReviewAnswer{}
		var hasComment int
		var choices string

		err = rows.Scan(&temp.ID, &temp.UserReviewer, &temp.UserTarget, &temp.ReviewID, &temp.AssignmentID,
			&temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Answer, &temp.Comment, &temp.Submitted,
			&hasComment, &choices, &temp.Weight)
		if err != nil {
			return result, err
		}

		temp.HasComment = hasComment == 1
		temp.Choices = strings.Split(choices, "|")

		result = append(result, &temp)
	}

	return result, err
}

// FetchForAssignment func
func (repo *ReviewAnswerRepository) FetchForAssignment(assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.assignment_id = ?"
	return fetchMany(repo, query, assignmentID)
}

// FetchForTarget func
func (repo *ReviewAnswerRepository) FetchForTarget(target, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_target = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, target, assignmentID)
	//return repo.FetchMany(query, target, assignmentID)
}

// FetchForReviewer func
func (repo *ReviewAnswerRepository) FetchForReviewer(reviewer, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, reviewer, assignmentID)
}

// FetchForReviewerAndTarget func
func (repo *ReviewAnswerRepository) FetchForReviewerAndTarget(reviewer, target, assignmentID int) ([]*model.ReviewAnswer, error) {
	query := "SELECT ur.id, ur.user_reviewer, ur.user_target, ur.review_id, ur.assignment_id, f.type, f.name, f.label, f.description, ur.answer, ur.comment, ur.submitted, f.hasComment, f.choices, f.weight FROM user_reviews AS ur INNER JOIN fields AS f ON ur.name = f.name WHERE ur.user_reviewer = ? AND ur.user_target = ? AND ur.assignment_id = ?"
	return fetchMany(repo, query, reviewer, target, assignmentID)
}

// Insert func
func (repo *ReviewAnswerRepository) Insert(answer model.ReviewAnswer) (int, error) {
	var id int64

	query := "INSERT INTO user_reviews (user_reviewer, user_target, review_id, assignment_id, type, name, label, answer, comment, submitted) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	tx, err := repo.db.Begin()
	if err != nil {
		return int(id), err
	}

	created := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	rows, err := tx.Exec(query, answer.UserReviewer, answer.UserTarget, answer.ReviewID, answer.AssignmentID,
		answer.Type, answer.Name, answer.Label, answer.Answer, answer.Comment, created)
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

// DeleteTarget func
func (repo *ReviewAnswerRepository) DeleteTarget(assignmentID, userID int) error {
	query := "DELETE FROM user_reviews WHERE assignment_id = ? AND user_target = ?"

	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}

	_, err = repo.db.Exec(query, assignmentID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return err
}
