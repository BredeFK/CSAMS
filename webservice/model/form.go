package model

import (
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/db"
	"github.com/JohanAanesen/NTNU-Bachelor-Management-System-For-CS-Assignments/webservice/shared/util"
	"time"
)

// Form TODO (Svein): comment
type Form struct {
	ID          int       `json:"id" db:"id"`
	Prefix      string    `json:"prefix" db:"prefix"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Created     time.Time `json:"created" db:"created"`
	Fields      []Field   `json:"fields"`
}

// Field TODO (Svein): comment
type Field struct {
	ID          int    `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`
	Name        string `json:"name" db:"name"`
	Label       string `json:"label" db:"label"`
	Description string `json:"description" db:"description"`
	Order       int    `json:"order" db:"priority"`
	Weight      int    `json:"weight" db:"weight"`
	Choices     string `json:"choices,omitempty" db:"choices"`
}

// Answer struct used for storing answers from users in forms
type Answer struct {
	ID    int
	Type  string
	Value string
}

// FormRepository ... TODO (Svein): comment
type FormRepository struct {
}

// Insert form to database
func (repo *FormRepository) Insert(form Form) (int, error) {

	// Get current Norwegian time in string format TODO time-norwegian
	date := util.ConvertTimeStampToString(util.GetTimeInNorwegian())

	// Insertions Query
	query := "INSERT INTO forms (prefix, name, description, created) VALUES (?, ?, ?, ?);"
	// Execute query with parameters
	rows, err := db.GetDB().Exec(query, form.Prefix, form.Name, form.Description, date)
	// Check for error
	if err != nil {
		return -1, err
	}

	// Get last inserted id from table
	id, err := rows.LastInsertId()
	// Check for error
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

// Get a single form based on the Primary Key, 'id'
func (repo *FormRepository) Get(id int) (Form, error) {
	var result = Form{}

	// Create query-string
	query := "SELECT id, prefix, name, description, created FROM forms WHERE id = ?"
	// Perform query
	rows, err := db.GetDB().Query(query, id)
	// Check for error
	if err != nil {
		return result, err
	}

	// Check if there is any rows
	if rows.Next() {
		// Scan
		err = rows.Scan(&result.ID, &result.Prefix, &result.Name, &result.Description, &result.Created)
		// Check for error
		if err != nil {
			return result, err
		}
	}
	// Create new query for getting all the fields
	query = "SELECT id, type, name, label, description, priority, weight, choices FROM fields WHERE form_id = ?"
	// Execute query
	rows, err = db.GetDB().Query(query, id)
	if err != nil {
		return result, err
	}

	// Loop through all rows
	for rows.Next() {
		var temp Field
		// Get values
		err = rows.Scan(&temp.ID, &temp.Type, &temp.Name, &temp.Label, &temp.Description, &temp.Order, &temp.Weight, &temp.Choices)
		if err != nil {
			return result, err
		}
		// Append field to slice in the result
		result.Fields = append(result.Fields, temp)
	}

	return result, err
}

// GetSubmissionFormFromAssignmentID get form from the assignment id key
func (repo *FormRepository) GetSubmissionFormFromAssignmentID(assignmentID int) (Form, error) {

	// Create query-string
	query := "SELECT f.form_id, f.id, f.type, f.name, f.label, f.description, f.priority, f.weight, f.choices FROM fields AS f WHERE f.form_id IN (SELECT s.form_id FROM submissions AS s WHERE id IN (SELECT a.submission_id FROM assignments AS a WHERE id=?)) ORDER BY f.priority"

	// Perform query
	rows, err := db.GetDB().Query(query, assignmentID)

	// Declare an empty Form
	form := Form{}

	// Check for error
	if err != nil {
		return form, err
	}

	for rows.Next() {
		var formID int
		var fieldID int
		var fieldType string
		var name string
		var label string
		var desc string
		var priority int
		var weight int
		var choices string

		// Scan
		err = rows.Scan(&formID, &fieldID, &fieldType, &name, &label, &desc, &priority, &weight, &choices)
		// Check for error
		if err != nil {
			return form, err
		}

		form.Fields = append(form.Fields, Field{
			ID:          formID,
			Type:        fieldType,
			Name:        name,
			Label:       label,
			Description: desc,
			Order:       priority,
			Weight:      weight,
			Choices:     choices,
		})
	}

	// TODO brede use sql.null<type>

	return form, nil
}

// GetReviewFormFromAssignmentID get review-form from the assignment id key
func (repo *FormRepository) GetReviewFormFromAssignmentID(assignmentID int) (Form, error) {

	// Create query-string
	query := "SELECT f.form_id, f.id, f.type, f.name, f.label, f.description, f.priority, f.weight, f.choices FROM fields AS f WHERE f.form_id IN (SELECT s.form_id FROM reviews AS s WHERE id IN (SELECT a.review_id FROM assignments AS a WHERE id=?)) ORDER BY f.priority"

	// Perform query
	rows, err := db.GetDB().Query(query, assignmentID)

	// Declare an empty Form
	form := Form{}

	// Check for error
	if err != nil {
		return form, err
	}

	for rows.Next() {
		var formID int
		var fieldID int
		var fieldType string
		var name string
		var label string
		var desc string
		var priority int
		var weight int
		var choices string

		// Scan
		err = rows.Scan(&formID, &fieldID, &fieldType, &name, &label, &desc, &priority, &weight, &choices)
		// Check for error
		if err != nil {
			return form, err
		}

		form.Fields = append(form.Fields, Field{
			ID:          formID,
			Type:        fieldType,
			Name:        name,
			Label:       label,
			Description: desc,
			Order:       priority,
			Weight:      weight,
			Choices:     choices,
		})
	}

	return form, err
}
