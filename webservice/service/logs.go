package service

import (
	"database/sql"
	"github.com/JohanAanesen/CSAMS/webservice/model"
	"github.com/JohanAanesen/CSAMS/webservice/repository"
)

// LogsService struct
type LogsService struct {
	logsRepo *repository.LogsRepository
}

// NewLogsService func
func NewLogsService(db *sql.DB) *LogsService {
	return &LogsService{
		logsRepo: repository.NewLogsRepository(db),
	}
}

// FetchAll fetches all logs
func (s *LogsService) FetchAll() ([]*model.Logs, error) {
	return s.logsRepo.FetchAll()
}

// InsertNewUser inserts a new user log
func (s *LogsService) InsertNewUser(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.NewUser,
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangeEmail inserts a change email log
func (s *LogsService) InsertChangeEmail(userID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangeEmail,
	}

	// Add oldValue to struct
	logx.OldValue = sql.NullString{
		String: oldValue,
		Valid:  oldValue != "",
	}

	// Add newValue to struct
	logx.NewValue = sql.NullString{
		String: newValue,
		Valid:  newValue != "",
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangeFAQ inserts a change FAQ log
func (s *LogsService) InsertChangeFAQ(userID int, oldValue string, newValue string) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateFAQ,
	}

	// Add oldValue to struct
	logx.OldValue = sql.NullString{
		String: oldValue,
		Valid:  oldValue != "",
	}

	// Add newValue to struct
	logx.NewValue = sql.NullString{
		String: newValue,
		Valid:  newValue != "",
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangePassword inserts a change password log
func (s *LogsService) InsertChangePassword(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangePassword,
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangePasswordEmail inserts a change password log
func (s *LogsService) InsertChangePasswordEmail(userID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.ChangePasswordEmail,
	}

	return s.logsRepo.Insert(logx)
}

// InsertAssignment inserts a change password log
func (s *LogsService) InsertAssignment(userID int, assignmentID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.DeliveredSubmission,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertChangeAssignment inserts a change password log
func (s *LogsService) InsertChangeAssignment(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.UpdateSubmission,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertUpdateAssignment inserts a change password log
func (s *LogsService) InsertUpdateAssignment(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteAssignment inserts a change password log
func (s *LogsService) InsertDeleteAssignment(userID int, assignmentID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteAssignment,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertFinishedOnePeerReview is for when one user has finished peer reviewing another users submission
func (s *LogsService) InsertFinishedOnePeerReview(userID int, assignmentID int, submissionID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.FinishedOnePeerReview,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}
	return s.logsRepo.Insert(logx)
}

// InsertUpdateOnePeerReview is for when one user has updated peer review
func (s *LogsService) InsertUpdateOnePeerReview(userID int, assignmentID int, submissionID int, affectedUserID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.UpdateOnePeerReview,
	}

	// Add assignmentID to struct
	logx.AssignmentID = sql.NullInt64{
		Int64: int64(assignmentID),
		Valid: assignmentID != 0,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	// Add affectedUserID to struct
	logx.AffectedUserID = sql.NullInt64{
		Int64: int64(affectedUserID),
		Valid: affectedUserID != 0,
	}
	return s.logsRepo.Insert(logx)
}

// InsertCourse inserts a new course log
func (s *LogsService) InsertCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreatedCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertJoinCourse inserts a new join course log
func (s *LogsService) InsertJoinCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.JoinedCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertLeftCourse inserts a new join course log
func (s *LogsService) InsertLeftCourse(userID int, courseID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.LeftCourse,
	}

	// Add courseID to struct
	logx.CourseID = sql.NullInt64{
		Int64: int64(courseID),
		Valid: courseID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertSubmissionForm inserts a new submission form
func (s *LogsService) InsertSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertUpdateSubmissionForm inserts a new submission form
func (s *LogsService) InsertUpdateSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteSubmissionForm inserts a new submission form
func (s *LogsService) InsertDeleteSubmissionForm(userID int, submissionID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteSubmissionForm,
	}

	// Add submissionID to struct
	logx.SubmissionID = sql.NullInt64{
		Int64: int64(submissionID),
		Valid: submissionID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertReviewForm inserts a new review form
func (s *LogsService) InsertReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminCreateReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertUpdateReviewForm inserts a new review form
func (s *LogsService) InsertUpdateReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminUpdateReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}

// InsertDeleteReviewForm inserts a new review form
func (s *LogsService) InsertDeleteReviewForm(userID int, reviewID int) error {

	// Save log in struct
	// logx since log is already an package
	logx := model.Logs{
		UserID:   userID,
		Activity: model.AdminDeleteReviewForm,
	}

	// Add reviewID to struct
	logx.ReviewID = sql.NullInt64{
		Int64: int64(reviewID),
		Valid: reviewID != 0,
	}

	return s.logsRepo.Insert(logx)
}
