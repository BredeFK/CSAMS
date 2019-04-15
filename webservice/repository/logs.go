package repository

import (
	"database/sql"
	"errors"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/shared/util"
)

// LogsRepository struct
type LogsRepository struct {
	db *sql.DB
}

// NewLogsRepository func
func NewLogsRepository(db *sql.DB) *LogsRepository {
	return &LogsRepository{
		db: db,
	}
}

// FetchAll fetches all logs
func (repo *LogsRepository) FetchAll() ([]*model.Logs, error) {
	result := make([]*model.Logs, 0)

	// Query to be executed
	query := "SELECT `id`, `user_id`, `timestamp`, `activity`, `assignment_id`, `course_id`, `submission_id`, `review_id`, `old_value`, `new_value`, `affected_user_id` FROM `logs`"

	// Run query
	rows, err := repo.db.Query(query)
	if err != nil {
		return result, err
	}

	// Close rows
	defer rows.Close()

	// Go through all rows
	for rows.Next() {
		temp := model.Logs{}

		// Add to temporary struct
		err = rows.Scan(&temp.ID, &temp.UserID, &temp.Timestamp, &temp.Activity, &temp.AssignmentID, &temp.CourseID,
			&temp.SubmissionID, &temp.ReviewID, &temp.OldValue, &temp.NewValue, &temp.AffectedUserID)
		if err != nil {
			return result, err
		}

		// Append to result array
		result = append(result, &temp)
	}

	return result, err
}

// Insert inserts all types of logs
func (repo *LogsRepository) Insert(logx model.Logs) error {

	// Different sql queries to different log types belows
	var err error

	tx, err := repo.db.Begin() //start transaction
	if err != nil {
		return err
	}

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInCorrectTimeZone())

	// Check what needs to be logged
	switch logx.Activity {
	case model.NewUser:
		err = newUser(tx, logx, date)
	case model.ChangeEmail:
		err = changeEmailUpdateFaq(tx, logx, date)
	case model.ChangePassword:
		err = changePassword(tx, logx, date)
	case model.ChangePasswordEmail:
		err = changePassword(tx, logx, date)
	case model.CreateSubmission:
		err = manageAssignmentAndFinishedPeer(tx, logx, date)
	case model.UpdateSubmission:
		err = manageAssignmentAndFinishedPeer(tx, logx, date)
	case model.DeleteSubmission:
		err = manageAssignmentAndFinishedPeer(tx, logx, date)
	case model.FinishedOnePeerReview:
		err = finishedOnePeerReview(tx, logx, date)
	case model.UpdateOnePeerReview:
		err = finishedOnePeerReview(tx, logx, date)
	case model.JoinedCourse:
		err = manageCourse(tx, logx, date)
	case model.LeftCourse:
		err = manageCourse(tx, logx, date)
	case model.AdminUpdateFAQ:
		err = changeEmailUpdateFaq(tx, logx, date)
	case model.AdminCreatAssignment:
		err = createAssignment(tx, logx, date)
	case model.AdminDeleteAssignment:
		err = manageAssignmentAndFinishedPeer(tx, logx, date)
	case model.AdminUpdateAssignment:
		err = manageAssignmentAndFinishedPeer(tx, logx, date)
	case model.AdminCreatedCourse:
		err = manageCourse(tx, logx, date)
	case model.AdminUpdateCourse:
		err = manageCourse(tx, logx, date)
	case model.AdminDeleteCourse:
		err = manageCourse(tx, logx, date)
	case model.AdminCreateSubmissionForm:
		err = manageSubmissionForm(tx, logx, date)
	case model.AdminUpdateSubmissionForm:
		err = manageSubmissionForm(tx, logx, date)
	case model.AdminDeleteSubmissionForm:
		err = manageSubmissionForm(tx, logx, date)
	case model.AdminCreateReviewForm:
		err = manageReviewForm(tx, logx, date)
	case model.AdminUpdateReviewForm:
		err = manageReviewForm(tx, logx, date)
	case model.AdminDeleteReviewForm:
		err = manageReviewForm(tx, logx, date)
	case model.AdminEmailCourseStudents:
		err = emailCourseStudents(tx, logx, date)
	case model.AdminRemoveUserFromCourse:
		err = adminRemoveUserFromCourse(tx, logx, date)
	case model.AdminChangeStudentPassword:
		err = adminChangeUserPassword(tx, logx, date)
	case model.AdminCreateSubmissionForUser:
		err = adminManageSubmissionForUser(tx, logx, date)
	case model.AdminUpdateSubmissionForUser:
		err = adminManageSubmissionForUser(tx, logx, date)
	case model.AdminDeleteSubmissionForUser:
		err = adminManageSubmissionForUser(tx, logx, date)
	default:
		var ErrLogActivityNotFound = errors.New("error: log activity not found")
		return ErrLogActivityNotFound
	}

	// Handle possible error
	if err != nil {
		tx.Rollback() //quit transaction if error
		return err
	}

	err = tx.Commit() //finish transaction
	if err != nil {
		return err
	}

	// Nothing went wrong -> return nil errors
	return nil
}

// newUser query for inserting new user log
func newUser(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", logx.UserID, date, logx.Activity)

	return err
}

// changeEmailUpdateFaq query for inserting change email or update faq log
func changeEmailUpdateFaq(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`, `Activity`, `old_value`, `new_value`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.OldValue, logx.NewValue)

	return err
}

// changePassword query for inserting change password log
func changePassword(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`, `Activity`) "+
		"VALUES (?, ?, ?)", logx.UserID, date, logx.Activity)

	return err
}

// manageAssignmentAndFinishedPeer query for inserting delete/update/deliver assignment and one review done and all reviews on one users review done log
func manageAssignmentAndFinishedPeer(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `assignment_id`, `submission_id`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentID, logx.SubmissionID)

	return err
}

// finishedOnePeerReview query for inserting when one user has review another users submission
func finishedOnePeerReview(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `assignment_id`, `submission_id`, `affected_user_id`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentID, logx.SubmissionID, logx.AffectedUserID)

	return err
}

// createAssignment query for inserting create assignment log
func createAssignment(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`, `Activity`, `assignment_id`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentID)

	return err
}

// manageCourse query for inserting join/create course log
func manageCourse(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `course_id`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.CourseID)

	return err
}

// manageSubmissionForm query for inserting create/update submission form
func manageSubmissionForm(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `submission_id`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.SubmissionID)

	return err
}

// manageReviewForm query for inserting create/update review form
func manageReviewForm(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `review_id`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.ReviewID)

	return err
}

// emailCourseStudents query for emailing students
func emailCourseStudents(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `course_id`, `new_value`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.CourseID, logx.NewValue)

	return err
}

// adminRemoveUserFromCourse
func adminRemoveUserFromCourse(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `course_id`, `affected_user_id`) "+
		"VALUES (?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.CourseID, logx.AffectedUserID)

	return err
}

// adminChangeUserPassword
func adminChangeUserPassword(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `affected_user_id`) "+
		"VALUES (?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AffectedUserID)

	return err
}

// adminManageSubmissionForUser
func adminManageSubmissionForUser(tx *sql.Tx, logx model.Logs, date string) error {
	_, err := tx.Query("INSERT INTO `logs` (`user_id`, `timestamp`,  `Activity`, `assignment_id`, `submission_id`, `affected_user_id`) "+
		"VALUES (?, ?, ?, ?, ?, ?)", logx.UserID, date, logx.Activity, logx.AssignmentID, logx.SubmissionID, logx.AffectedUserID)

	return err
}
